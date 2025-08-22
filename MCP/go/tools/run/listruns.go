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

func ListrunsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("name=%v", val))
		}
		if val, ok := args["runGroupId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("runGroupId=%v", val))
		}
		if val, ok := args["startingToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("startingToken=%v", val))
		}
		if val, ok := args["maxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxResults=%v", val))
		}
		if val, ok := args["status"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("status=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/run%s", cfg.BaseURL, queryString)
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
		var result models.ListRunsResponse
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

func CreateListrunsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_run",
		mcp.WithDescription("Retrieves a list of runs."),
		mcp.WithString("name", mcp.Description("Filter the list by run name.")),
		mcp.WithString("runGroupId", mcp.Description("Filter the list by run group ID.")),
		mcp.WithString("startingToken", mcp.Description("Specify the pagination token from a previous request to retrieve the next page of results.")),
		mcp.WithNumber("maxResults", mcp.Description("The maximum number of runs to return in one page of results.")),
		mcp.WithString("status", mcp.Description(" The status of a run. ")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ListrunsHandler(cfg),
	}
}
