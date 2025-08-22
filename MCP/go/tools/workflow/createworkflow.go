package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/amazon-omics/mcp-server/config"
	"github.com/amazon-omics/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func CreateworkflowHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/workflow", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.CreateWorkflowResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateCreateworkflowTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_workflow",
		mcp.WithDescription("Creates a workflow."),
		mcp.WithString("definitionZip", mcp.Description("Input parameter: A ZIP archive for the workflow.")),
		mcp.WithString("main", mcp.Description("Input parameter: The path of the main definition file for the workflow.")),
		mcp.WithString("name", mcp.Description("Input parameter: A name for the workflow.")),
		mcp.WithString("requestId", mcp.Required(), mcp.Description("Input parameter: To ensure that requests don't run multiple times, specify a unique ID for each request.")),
		mcp.WithString("engine", mcp.Description("Input parameter: An engine for the workflow.")),
		mcp.WithString("accelerators", mcp.Description("Input parameter:  The computational accelerator specified to run the workflow. ")),
		mcp.WithObject("tags", mcp.Description("Input parameter: Tags for the workflow.")),
		mcp.WithString("description", mcp.Description("Input parameter: A description for the workflow.")),
		mcp.WithObject("parameterTemplate", mcp.Description("Input parameter: A parameter template for the workflow.")),
		mcp.WithNumber("storageCapacity", mcp.Description("Input parameter: A storage capacity for the workflow in gigabytes.")),
		mcp.WithString("definitionUri", mcp.Description("Input parameter: The URI of a definition for the workflow.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreateworkflowHandler(cfg),
	}
}
