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

func CompletemultipartreadsetuploadHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		uploadIdVal, ok := args["uploadId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: uploadId"), nil
		}
		uploadId, ok := uploadIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: uploadId"), nil
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
		url := fmt.Sprintf("%s/sequencestore/%s/upload/%s/complete", cfg.BaseURL, sequenceStoreId, uploadId)
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
		var result models.CompleteMultipartReadSetUploadResponse
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

func CreateCompletemultipartreadsetuploadTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_sequencestore_sequenceStoreId_upload_uploadId_complete",
		mcp.WithDescription(" Concludes a multipart upload once you have uploaded all the components. "),
		mcp.WithString("sequenceStoreId", mcp.Required(), mcp.Description(" The sequence store ID for the store involved in the multipart upload. ")),
		mcp.WithString("uploadId", mcp.Required(), mcp.Description(" The ID for the multipart upload. ")),
		mcp.WithArray("parts", mcp.Required(), mcp.Description("Input parameter:  The individual uploads or parts of a multipart upload. ")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CompletemultipartreadsetuploadHandler(cfg),
	}
}
