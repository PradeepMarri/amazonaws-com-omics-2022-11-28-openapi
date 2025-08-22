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

func StartreadsetexportjobHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		sequenceStoreIdVal, ok := args["sequenceStoreId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: sequenceStoreId"), nil
		}
		sequenceStoreId, ok := sequenceStoreIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: sequenceStoreId"), nil
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
		url := fmt.Sprintf("%s/sequencestore/%s/exportjob", cfg.BaseURL, sequenceStoreId)
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
		var result models.StartReadSetExportJobResponse
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

func CreateStartreadsetexportjobTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_sequencestore_sequenceStoreId_exportjob",
		mcp.WithDescription("Exports a read set to Amazon S3."),
		mcp.WithString("sequenceStoreId", mcp.Required(), mcp.Description("The read set's sequence store ID.")),
		mcp.WithString("destination", mcp.Required(), mcp.Description("Input parameter: A location for exported files in Amazon S3.")),
		mcp.WithString("roleArn", mcp.Required(), mcp.Description("Input parameter: A service role for the job.")),
		mcp.WithArray("sources", mcp.Required(), mcp.Description("Input parameter: The job's source files.")),
		mcp.WithString("clientToken", mcp.Description("Input parameter: To ensure that jobs don't run multiple times, specify a unique token for each job.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    StartreadsetexportjobHandler(cfg),
	}
}
