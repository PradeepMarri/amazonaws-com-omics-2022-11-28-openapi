package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/amazon-omics/mcp-server/config"
	"github.com/amazon-omics/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetreadsetmetadataHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		idVal, ok := args["id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: id"), nil
		}
		id, ok := idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: id"), nil
		}
		sequenceStoreIdVal, ok := args["sequenceStoreId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: sequenceStoreId"), nil
		}
		sequenceStoreId, ok := sequenceStoreIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: sequenceStoreId"), nil
		}
		url := fmt.Sprintf("%s/sequencestore/%s/readset/%s/metadata", cfg.BaseURL, id, sequenceStoreId)
		req, err := http.NewRequest("GET", url, nil)
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
		var result models.GetReadSetMetadataResponse
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

func CreateGetreadsetmetadataTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_sequencestore_sequenceStoreId_readset_id_metadata",
		mcp.WithDescription("Gets details about a read set."),
		mcp.WithString("id", mcp.Required(), mcp.Description("The read set's ID.")),
		mcp.WithString("sequenceStoreId", mcp.Required(), mcp.Description("The read set's sequence store ID.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetreadsetmetadataHandler(cfg),
	}
}
