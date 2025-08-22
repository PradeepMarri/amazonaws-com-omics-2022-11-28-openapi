package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/amazon-omics/mcp-server/config"
	"github.com/amazon-omics/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func ListmultipartreadsetuploadsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		queryParams := make([]string, 0)
		if val, ok := args["maxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxResults=%v", val))
		}
		if val, ok := args["nextToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("nextToken=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/sequencestore/%s/uploads%s", cfg.BaseURL, sequenceStoreId, queryString)
		req, err := http.NewRequest("POST", url, nil)
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
		var result models.ListMultipartReadSetUploadsResponse
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

func CreateListmultipartreadsetuploadsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_sequencestore_sequenceStoreId_uploads",
		mcp.WithDescription(" Lists all multipart read set uploads and their statuses. "),
		mcp.WithString("sequenceStoreId", mcp.Required(), mcp.Description(" The Sequence Store ID used for the multipart uploads. ")),
		mcp.WithNumber("maxResults", mcp.Description(" The maximum number of multipart uploads returned in a page. ")),
		mcp.WithString("nextToken", mcp.Description(" Next token returned in the response of a previous ListMultipartReadSetUploads call. Used to get the next page of results. ")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ListmultipartreadsetuploadsHandler(cfg),
	}
}
