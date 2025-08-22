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

func CreatemultipartreadsetuploadHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/sequencestore/%s/upload", cfg.BaseURL, sequenceStoreId)
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
		var result models.CreateMultipartReadSetUploadResponse
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

func CreateCreatemultipartreadsetuploadTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_sequencestore_sequenceStoreId_upload",
		mcp.WithDescription(" Begins a multipart read set upload. "),
		mcp.WithString("sequenceStoreId", mcp.Required(), mcp.Description(" The sequence store ID for the store that is the destination of the multipart uploads. ")),
		mcp.WithString("description", mcp.Description("Input parameter:  The description of the read set. ")),
		mcp.WithString("name", mcp.Required(), mcp.Description("Input parameter:  The name of the read set. ")),
		mcp.WithString("sampleId", mcp.Required(), mcp.Description("Input parameter:  The source's sample ID. ")),
		mcp.WithString("clientToken", mcp.Description("Input parameter:  An idempotency token that can be used to avoid triggering multiple multipart uploads. ")),
		mcp.WithString("generatedFrom", mcp.Description("Input parameter:  Where the source originated. ")),
		mcp.WithString("subjectId", mcp.Required(), mcp.Description("Input parameter:  The source's subject ID. ")),
		mcp.WithString("sourceFileType", mcp.Required(), mcp.Description("Input parameter:  The type of file being uploaded. ")),
		mcp.WithString("referenceArn", mcp.Required(), mcp.Description("Input parameter:  The ARN of the reference. ")),
		mcp.WithObject("tags", mcp.Description("Input parameter:  Any tags to add to the read set. ")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreatemultipartreadsetuploadHandler(cfg),
	}
}
