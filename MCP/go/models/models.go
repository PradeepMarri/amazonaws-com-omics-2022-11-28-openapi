package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ListVariantImportJobsResponse represents the ListVariantImportJobsResponse schema from the OpenAPI specification
type ListVariantImportJobsResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Variantimportjobs interface{} `json:"variantImportJobs,omitempty"`
}

// CreateReferenceStoreResponse represents the CreateReferenceStoreResponse schema from the OpenAPI specification
type CreateReferenceStoreResponse struct {
	Description interface{} `json:"description,omitempty"`
	Id interface{} `json:"id"`
	Name interface{} `json:"name,omitempty"`
	Sseconfig interface{} `json:"sseConfig,omitempty"`
	Arn interface{} `json:"arn"`
	Creationtime interface{} `json:"creationTime"`
}

// ListReferencesResponse represents the ListReferencesResponse schema from the OpenAPI specification
type ListReferencesResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	References interface{} `json:"references"`
}

// AbortMultipartReadSetUploadRequest represents the AbortMultipartReadSetUploadRequest schema from the OpenAPI specification
type AbortMultipartReadSetUploadRequest struct {
}

// ListWorkflowsRequest represents the ListWorkflowsRequest schema from the OpenAPI specification
type ListWorkflowsRequest struct {
}

// GetReferenceStoreRequest represents the GetReferenceStoreRequest schema from the OpenAPI specification
type GetReferenceStoreRequest struct {
}

// GetVariantStoreRequest represents the GetVariantStoreRequest schema from the OpenAPI specification
type GetVariantStoreRequest struct {
}

// CreateAnnotationStoreResponse represents the CreateAnnotationStoreResponse schema from the OpenAPI specification
type CreateAnnotationStoreResponse struct {
	Storeoptions interface{} `json:"storeOptions,omitempty"`
	Creationtime interface{} `json:"creationTime"`
	Id interface{} `json:"id"`
	Name interface{} `json:"name"`
	Reference interface{} `json:"reference,omitempty"`
	Status interface{} `json:"status"`
	Storeformat interface{} `json:"storeFormat,omitempty"`
}

// ListVariantImportJobsFilter represents the ListVariantImportJobsFilter schema from the OpenAPI specification
type ListVariantImportJobsFilter struct {
	Status interface{} `json:"status,omitempty"`
	Storename interface{} `json:"storeName,omitempty"`
}

// CancelAnnotationImportRequest represents the CancelAnnotationImportRequest schema from the OpenAPI specification
type CancelAnnotationImportRequest struct {
}

// DeleteSequenceStoreRequest represents the DeleteSequenceStoreRequest schema from the OpenAPI specification
type DeleteSequenceStoreRequest struct {
}

// ListWorkflowsResponse represents the ListWorkflowsResponse schema from the OpenAPI specification
type ListWorkflowsResponse struct {
	Items interface{} `json:"items,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// ListReferenceImportJobsResponse represents the ListReferenceImportJobsResponse schema from the OpenAPI specification
type ListReferenceImportJobsResponse struct {
	Importjobs interface{} `json:"importJobs,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// UpdateRunGroupRequest represents the UpdateRunGroupRequest schema from the OpenAPI specification
type UpdateRunGroupRequest struct {
	Maxgpus interface{} `json:"maxGpus,omitempty"`
	Maxruns interface{} `json:"maxRuns,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Maxcpus interface{} `json:"maxCpus,omitempty"`
	Maxduration interface{} `json:"maxDuration,omitempty"`
}

// ImportReferenceSourceItem represents the ImportReferenceSourceItem schema from the OpenAPI specification
type ImportReferenceSourceItem struct {
	Status interface{} `json:"status"`
	Statusmessage interface{} `json:"statusMessage,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Sourcefile interface{} `json:"sourceFile,omitempty"`
}

// FormatOptions represents the FormatOptions schema from the OpenAPI specification
type FormatOptions struct {
	Vcfoptions interface{} `json:"vcfOptions,omitempty"`
	Tsvoptions interface{} `json:"tsvOptions,omitempty"`
}

// ListVariantStoresRequest represents the ListVariantStoresRequest schema from the OpenAPI specification
type ListVariantStoresRequest struct {
	Ids interface{} `json:"ids,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

// VariantImportJobItem represents the VariantImportJobItem schema from the OpenAPI specification
type VariantImportJobItem struct {
	Annotationfields interface{} `json:"annotationFields,omitempty"`
	Rolearn interface{} `json:"roleArn"`
	Updatetime interface{} `json:"updateTime"`
	Completiontime interface{} `json:"completionTime,omitempty"`
	Creationtime interface{} `json:"creationTime"`
	Destinationname interface{} `json:"destinationName"`
	Id interface{} `json:"id"`
	Status interface{} `json:"status"`
	Runleftnormalization interface{} `json:"runLeftNormalization,omitempty"`
}

// GetAnnotationStoreRequest represents the GetAnnotationStoreRequest schema from the OpenAPI specification
type GetAnnotationStoreRequest struct {
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
}

// CreateRunGroupResponse represents the CreateRunGroupResponse schema from the OpenAPI specification
type CreateRunGroupResponse struct {
	Arn interface{} `json:"arn,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
}

// GetAnnotationImportResponse represents the GetAnnotationImportResponse schema from the OpenAPI specification
type GetAnnotationImportResponse struct {
	Runleftnormalization interface{} `json:"runLeftNormalization"`
	Status interface{} `json:"status"`
	Rolearn interface{} `json:"roleArn"`
	Statusmessage interface{} `json:"statusMessage"`
	Completiontime interface{} `json:"completionTime"`
	Creationtime interface{} `json:"creationTime"`
	Destinationname interface{} `json:"destinationName"`
	Items interface{} `json:"items"`
	Updatetime interface{} `json:"updateTime"`
	Annotationfields interface{} `json:"annotationFields,omitempty"`
	Formatoptions FormatOptions `json:"formatOptions"` // Formatting options for a file.
	Id interface{} `json:"id"`
}

// StartReadSetImportJobRequest represents the StartReadSetImportJobRequest schema from the OpenAPI specification
type StartReadSetImportJobRequest struct {
	Sources interface{} `json:"sources"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Rolearn interface{} `json:"roleArn"`
}

// GetReadSetRequest represents the GetReadSetRequest schema from the OpenAPI specification
type GetReadSetRequest struct {
}

// ReadSetFilter represents the ReadSetFilter schema from the OpenAPI specification
type ReadSetFilter struct {
	Status interface{} `json:"status,omitempty"`
	Referencearn interface{} `json:"referenceArn,omitempty"`
	Createdafter interface{} `json:"createdAfter,omitempty"`
	Generatedfrom interface{} `json:"generatedFrom,omitempty"`
	Subjectid interface{} `json:"subjectId,omitempty"`
	Createdbefore interface{} `json:"createdBefore,omitempty"`
	Creationtype interface{} `json:"creationType,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Sampleid interface{} `json:"sampleId,omitempty"`
}

// TsvOptions represents the TsvOptions schema from the OpenAPI specification
type TsvOptions struct {
	Readoptions interface{} `json:"readOptions,omitempty"`
}

// DeleteVariantStoreRequest represents the DeleteVariantStoreRequest schema from the OpenAPI specification
type DeleteVariantStoreRequest struct {
}

// GetReferenceMetadataRequest represents the GetReferenceMetadataRequest schema from the OpenAPI specification
type GetReferenceMetadataRequest struct {
}

// UpdateVariantStoreRequest represents the UpdateVariantStoreRequest schema from the OpenAPI specification
type UpdateVariantStoreRequest struct {
	Description interface{} `json:"description,omitempty"`
}

// ReadSetListItem represents the ReadSetListItem schema from the OpenAPI specification
type ReadSetListItem struct {
	Sampleid interface{} `json:"sampleId,omitempty"`
	Statusmessage interface{} `json:"statusMessage,omitempty"`
	Subjectid interface{} `json:"subjectId,omitempty"`
	Filetype interface{} `json:"fileType"`
	Id interface{} `json:"id"`
	Referencearn interface{} `json:"referenceArn,omitempty"`
	Sequenceinformation SequenceInformation `json:"sequenceInformation,omitempty"` // Details about a sequence.
	Status interface{} `json:"status"`
	Creationtime interface{} `json:"creationTime"`
	Name interface{} `json:"name,omitempty"`
	Arn interface{} `json:"arn"`
	Creationtype interface{} `json:"creationType,omitempty"`
	Sequencestoreid interface{} `json:"sequenceStoreId"`
	Description interface{} `json:"description,omitempty"`
}

// ListSequenceStoresResponse represents the ListSequenceStoresResponse schema from the OpenAPI specification
type ListSequenceStoresResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Sequencestores interface{} `json:"sequenceStores"`
}

// StartReadSetActivationJobResponse represents the StartReadSetActivationJobResponse schema from the OpenAPI specification
type StartReadSetActivationJobResponse struct {
	Id interface{} `json:"id"`
	Sequencestoreid interface{} `json:"sequenceStoreId"`
	Status interface{} `json:"status"`
	Creationtime interface{} `json:"creationTime"`
}

// BatchDeleteReadSetResponse represents the BatchDeleteReadSetResponse schema from the OpenAPI specification
type BatchDeleteReadSetResponse struct {
	Errors interface{} `json:"errors,omitempty"`
}

// AnnotationImportJobItem represents the AnnotationImportJobItem schema from the OpenAPI specification
type AnnotationImportJobItem struct {
	Status interface{} `json:"status"`
	Completiontime interface{} `json:"completionTime,omitempty"`
	Creationtime interface{} `json:"creationTime"`
	Id interface{} `json:"id"`
	Runleftnormalization interface{} `json:"runLeftNormalization,omitempty"`
	Annotationfields interface{} `json:"annotationFields,omitempty"`
	Rolearn interface{} `json:"roleArn"`
	Destinationname interface{} `json:"destinationName"`
	Updatetime interface{} `json:"updateTime"`
}

// ListVariantImportJobsRequest represents the ListVariantImportJobsRequest schema from the OpenAPI specification
type ListVariantImportJobsRequest struct {
	Filter interface{} `json:"filter,omitempty"`
	Ids interface{} `json:"ids,omitempty"`
}

// TagMap represents the TagMap schema from the OpenAPI specification
type TagMap struct {
}

// GetVariantImportResponse represents the GetVariantImportResponse schema from the OpenAPI specification
type GetVariantImportResponse struct {
	Completiontime interface{} `json:"completionTime,omitempty"`
	Rolearn interface{} `json:"roleArn"`
	Annotationfields interface{} `json:"annotationFields,omitempty"`
	Destinationname interface{} `json:"destinationName"`
	Runleftnormalization interface{} `json:"runLeftNormalization"`
	Status interface{} `json:"status"`
	Statusmessage interface{} `json:"statusMessage"`
	Updatetime interface{} `json:"updateTime"`
	Creationtime interface{} `json:"creationTime"`
	Id interface{} `json:"id"`
	Items interface{} `json:"items"`
}

// GetReadSetExportJobResponse represents the GetReadSetExportJobResponse schema from the OpenAPI specification
type GetReadSetExportJobResponse struct {
	Status interface{} `json:"status"`
	Statusmessage interface{} `json:"statusMessage,omitempty"`
	Completiontime interface{} `json:"completionTime,omitempty"`
	Creationtime interface{} `json:"creationTime"`
	Destination interface{} `json:"destination"`
	Id interface{} `json:"id"`
	Readsets interface{} `json:"readSets,omitempty"`
	Sequencestoreid interface{} `json:"sequenceStoreId"`
}

// CreateAnnotationStoreRequest represents the CreateAnnotationStoreRequest schema from the OpenAPI specification
type CreateAnnotationStoreRequest struct {
	Description interface{} `json:"description,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Reference interface{} `json:"reference,omitempty"`
	Sseconfig interface{} `json:"sseConfig,omitempty"`
	Storeformat interface{} `json:"storeFormat"`
	Storeoptions interface{} `json:"storeOptions,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
}

// StartReadSetExportJobResponse represents the StartReadSetExportJobResponse schema from the OpenAPI specification
type StartReadSetExportJobResponse struct {
	Id interface{} `json:"id"`
	Sequencestoreid interface{} `json:"sequenceStoreId"`
	Status interface{} `json:"status"`
	Creationtime interface{} `json:"creationTime"`
	Destination interface{} `json:"destination"`
}

// ReferenceStoreDetail represents the ReferenceStoreDetail schema from the OpenAPI specification
type ReferenceStoreDetail struct {
	Description interface{} `json:"description,omitempty"`
	Id interface{} `json:"id"`
	Name interface{} `json:"name,omitempty"`
	Sseconfig interface{} `json:"sseConfig,omitempty"`
	Arn interface{} `json:"arn"`
	Creationtime interface{} `json:"creationTime"`
}

// ListReadSetExportJobsResponse represents the ListReadSetExportJobsResponse schema from the OpenAPI specification
type ListReadSetExportJobsResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Exportjobs interface{} `json:"exportJobs,omitempty"`
}

// GetAnnotationStoreResponse represents the GetAnnotationStoreResponse schema from the OpenAPI specification
type GetAnnotationStoreResponse struct {
	Tags interface{} `json:"tags"`
	Updatetime interface{} `json:"updateTime"`
	Reference interface{} `json:"reference"`
	Statusmessage interface{} `json:"statusMessage"`
	Creationtime interface{} `json:"creationTime"`
	Sseconfig interface{} `json:"sseConfig"`
	Status interface{} `json:"status"`
	Storeformat interface{} `json:"storeFormat,omitempty"`
	Description interface{} `json:"description"`
	Id interface{} `json:"id"`
	Name interface{} `json:"name"`
	Storeoptions interface{} `json:"storeOptions,omitempty"`
	Storearn interface{} `json:"storeArn"`
	Storesizebytes interface{} `json:"storeSizeBytes"`
}

// CreateReferenceStoreRequest represents the CreateReferenceStoreRequest schema from the OpenAPI specification
type CreateReferenceStoreRequest struct {
	Sseconfig interface{} `json:"sseConfig,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Name interface{} `json:"name"`
}

// TagResourceRequestTagsMap represents the TagResourceRequestTagsMap schema from the OpenAPI specification
type TagResourceRequestTagsMap struct {
}

// StartVariantImportResponse represents the StartVariantImportResponse schema from the OpenAPI specification
type StartVariantImportResponse struct {
	Jobid interface{} `json:"jobId"`
}

// StartReadSetActivationJobRequest represents the StartReadSetActivationJobRequest schema from the OpenAPI specification
type StartReadSetActivationJobRequest struct {
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Sources interface{} `json:"sources"`
}

// GetWorkflowRequest represents the GetWorkflowRequest schema from the OpenAPI specification
type GetWorkflowRequest struct {
}

// ExportReadSetDetail represents the ExportReadSetDetail schema from the OpenAPI specification
type ExportReadSetDetail struct {
	Id interface{} `json:"id"`
	Status interface{} `json:"status"`
	Statusmessage interface{} `json:"statusMessage,omitempty"`
}

// ListReadSetUploadPartsRequest represents the ListReadSetUploadPartsRequest schema from the OpenAPI specification
type ListReadSetUploadPartsRequest struct {
	Filter interface{} `json:"filter,omitempty"`
	Partsource interface{} `json:"partSource"`
}

// ListReferencesRequest represents the ListReferencesRequest schema from the OpenAPI specification
type ListReferencesRequest struct {
	Filter interface{} `json:"filter,omitempty"`
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// CancelAnnotationImportResponse represents the CancelAnnotationImportResponse schema from the OpenAPI specification
type CancelAnnotationImportResponse struct {
}

// GetRunGroupResponse represents the GetRunGroupResponse schema from the OpenAPI specification
type GetRunGroupResponse struct {
	Tags interface{} `json:"tags,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Maxcpus interface{} `json:"maxCpus,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Creationtime interface{} `json:"creationTime,omitempty"`
	Maxduration interface{} `json:"maxDuration,omitempty"`
	Maxgpus interface{} `json:"maxGpus,omitempty"`
	Maxruns interface{} `json:"maxRuns,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// GetReadSetExportJobRequest represents the GetReadSetExportJobRequest schema from the OpenAPI specification
type GetReadSetExportJobRequest struct {
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Tags interface{} `json:"tags"`
}

// CompleteMultipartReadSetUploadResponse represents the CompleteMultipartReadSetUploadResponse schema from the OpenAPI specification
type CompleteMultipartReadSetUploadResponse struct {
	Readsetid interface{} `json:"readSetId"`
}

// GetWorkflowResponse represents the GetWorkflowResponse schema from the OpenAPI specification
type GetWorkflowResponse struct {
	Engine interface{} `json:"engine,omitempty"`
	TypeField interface{} `json:"type,omitempty"`
	Parametertemplate interface{} `json:"parameterTemplate,omitempty"`
	Storagecapacity interface{} `json:"storageCapacity,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Main interface{} `json:"main,omitempty"`
	Accelerators interface{} `json:"accelerators,omitempty"`
	Creationtime interface{} `json:"creationTime,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Definition interface{} `json:"definition,omitempty"`
	Digest interface{} `json:"digest,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Statusmessage interface{} `json:"statusMessage,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
}

// ImportReadSetFilter represents the ImportReadSetFilter schema from the OpenAPI specification
type ImportReadSetFilter struct {
	Createdafter interface{} `json:"createdAfter,omitempty"`
	Createdbefore interface{} `json:"createdBefore,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// ReferenceItem represents the ReferenceItem schema from the OpenAPI specification
type ReferenceItem struct {
	Referencearn interface{} `json:"referenceArn,omitempty"`
}

// ActivateReadSetJobItem represents the ActivateReadSetJobItem schema from the OpenAPI specification
type ActivateReadSetJobItem struct {
	Sequencestoreid interface{} `json:"sequenceStoreId"`
	Status interface{} `json:"status"`
	Completiontime interface{} `json:"completionTime,omitempty"`
	Creationtime interface{} `json:"creationTime"`
	Id interface{} `json:"id"`
}

// CreateMultipartReadSetUploadResponse represents the CreateMultipartReadSetUploadResponse schema from the OpenAPI specification
type CreateMultipartReadSetUploadResponse struct {
	Generatedfrom interface{} `json:"generatedFrom,omitempty"`
	Sequencestoreid interface{} `json:"sequenceStoreId"`
	Sourcefiletype interface{} `json:"sourceFileType"`
	Subjectid interface{} `json:"subjectId"`
	Tags interface{} `json:"tags,omitempty"`
	Uploadid interface{} `json:"uploadId"`
	Description interface{} `json:"description,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Sampleid interface{} `json:"sampleId"`
	Creationtime interface{} `json:"creationTime"`
	Referencearn interface{} `json:"referenceArn"`
}

// StartReferenceImportJobSourceItem represents the StartReferenceImportJobSourceItem schema from the OpenAPI specification
type StartReferenceImportJobSourceItem struct {
	Tags interface{} `json:"tags,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Name interface{} `json:"name"`
	Sourcefile interface{} `json:"sourceFile"`
}

// RunParameters represents the RunParameters schema from the OpenAPI specification
type RunParameters struct {
}

// ListReferenceStoresResponse represents the ListReferenceStoresResponse schema from the OpenAPI specification
type ListReferenceStoresResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Referencestores interface{} `json:"referenceStores"`
}

// ListAnnotationImportJobsFilter represents the ListAnnotationImportJobsFilter schema from the OpenAPI specification
type ListAnnotationImportJobsFilter struct {
	Status interface{} `json:"status,omitempty"`
	Storename interface{} `json:"storeName,omitempty"`
}

// DeleteAnnotationStoreRequest represents the DeleteAnnotationStoreRequest schema from the OpenAPI specification
type DeleteAnnotationStoreRequest struct {
}

// GetReadSetResponse represents the GetReadSetResponse schema from the OpenAPI specification
type GetReadSetResponse struct {
	Payload interface{} `json:"payload,omitempty"`
}

// RunResourceDigests represents the RunResourceDigests schema from the OpenAPI specification
type RunResourceDigests struct {
}

// CreateWorkflowResponse represents the CreateWorkflowResponse schema from the OpenAPI specification
type CreateWorkflowResponse struct {
	Id interface{} `json:"id,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
}

// ListReferenceStoresRequest represents the ListReferenceStoresRequest schema from the OpenAPI specification
type ListReferenceStoresRequest struct {
	Filter interface{} `json:"filter,omitempty"`
}

// CancelVariantImportRequest represents the CancelVariantImportRequest schema from the OpenAPI specification
type CancelVariantImportRequest struct {
}

// AbortMultipartReadSetUploadResponse represents the AbortMultipartReadSetUploadResponse schema from the OpenAPI specification
type AbortMultipartReadSetUploadResponse struct {
}

// CompleteMultipartReadSetUploadRequest represents the CompleteMultipartReadSetUploadRequest schema from the OpenAPI specification
type CompleteMultipartReadSetUploadRequest struct {
	Parts interface{} `json:"parts"`
}

// ActivateReadSetSourceItem represents the ActivateReadSetSourceItem schema from the OpenAPI specification
type ActivateReadSetSourceItem struct {
	Statusmessage interface{} `json:"statusMessage,omitempty"`
	Readsetid interface{} `json:"readSetId"`
	Status interface{} `json:"status"`
}

// AnnotationImportItemDetail represents the AnnotationImportItemDetail schema from the OpenAPI specification
type AnnotationImportItemDetail struct {
	Jobstatus interface{} `json:"jobStatus"`
	Source interface{} `json:"source"`
}

// ListReadSetExportJobsRequest represents the ListReadSetExportJobsRequest schema from the OpenAPI specification
type ListReadSetExportJobsRequest struct {
	Filter interface{} `json:"filter,omitempty"`
}

// CreateSequenceStoreRequest represents the CreateSequenceStoreRequest schema from the OpenAPI specification
type CreateSequenceStoreRequest struct {
	Description interface{} `json:"description,omitempty"`
	Fallbacklocation interface{} `json:"fallbackLocation,omitempty"`
	Name interface{} `json:"name"`
	Sseconfig interface{} `json:"sseConfig,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
}

// GetReferenceMetadataResponse represents the GetReferenceMetadataResponse schema from the OpenAPI specification
type GetReferenceMetadataResponse struct {
	Description interface{} `json:"description,omitempty"`
	Files interface{} `json:"files,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Referencestoreid interface{} `json:"referenceStoreId"`
	Id interface{} `json:"id"`
	Md5 interface{} `json:"md5"`
	Status interface{} `json:"status,omitempty"`
	Creationtime interface{} `json:"creationTime"`
	Updatetime interface{} `json:"updateTime"`
	Arn interface{} `json:"arn"`
}

// GetReferenceRequest represents the GetReferenceRequest schema from the OpenAPI specification
type GetReferenceRequest struct {
}

// ImportReadSetJobItem represents the ImportReadSetJobItem schema from the OpenAPI specification
type ImportReadSetJobItem struct {
	Completiontime interface{} `json:"completionTime,omitempty"`
	Creationtime interface{} `json:"creationTime"`
	Id interface{} `json:"id"`
	Rolearn interface{} `json:"roleArn"`
	Sequencestoreid interface{} `json:"sequenceStoreId"`
	Status interface{} `json:"status"`
}

// ListVariantStoresFilter represents the ListVariantStoresFilter schema from the OpenAPI specification
type ListVariantStoresFilter struct {
	Status interface{} `json:"status,omitempty"`
}

// StartReadSetActivationJobSourceItem represents the StartReadSetActivationJobSourceItem schema from the OpenAPI specification
type StartReadSetActivationJobSourceItem struct {
	Readsetid interface{} `json:"readSetId"`
}

// StoreOptions represents the StoreOptions schema from the OpenAPI specification
type StoreOptions struct {
	Tsvstoreoptions interface{} `json:"tsvStoreOptions,omitempty"`
}

// ListReadSetsResponse represents the ListReadSetsResponse schema from the OpenAPI specification
type ListReadSetsResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Readsets interface{} `json:"readSets"`
}

// ExportReadSetJobDetail represents the ExportReadSetJobDetail schema from the OpenAPI specification
type ExportReadSetJobDetail struct {
	Destination interface{} `json:"destination"`
	Id interface{} `json:"id"`
	Sequencestoreid interface{} `json:"sequenceStoreId"`
	Status interface{} `json:"status"`
	Completiontime interface{} `json:"completionTime,omitempty"`
	Creationtime interface{} `json:"creationTime"`
}

// AnnotationFieldMap represents the AnnotationFieldMap schema from the OpenAPI specification
type AnnotationFieldMap struct {
}

// StartAnnotationImportRequest represents the StartAnnotationImportRequest schema from the OpenAPI specification
type StartAnnotationImportRequest struct {
	Annotationfields interface{} `json:"annotationFields,omitempty"`
	Destinationname interface{} `json:"destinationName"`
	Formatoptions interface{} `json:"formatOptions,omitempty"`
	Items interface{} `json:"items"`
	Rolearn interface{} `json:"roleArn"`
	Runleftnormalization interface{} `json:"runLeftNormalization,omitempty"`
}

// VariantImportItemDetail represents the VariantImportItemDetail schema from the OpenAPI specification
type VariantImportItemDetail struct {
	Jobstatus interface{} `json:"jobStatus"`
	Source interface{} `json:"source"`
	Statusmessage interface{} `json:"statusMessage,omitempty"`
}

// ListReadSetsRequest represents the ListReadSetsRequest schema from the OpenAPI specification
type ListReadSetsRequest struct {
	Filter interface{} `json:"filter,omitempty"`
}

// CreateRunGroupRequest represents the CreateRunGroupRequest schema from the OpenAPI specification
type CreateRunGroupRequest struct {
	Maxcpus interface{} `json:"maxCpus,omitempty"`
	Maxduration interface{} `json:"maxDuration,omitempty"`
	Maxgpus interface{} `json:"maxGpus,omitempty"`
	Maxruns interface{} `json:"maxRuns,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Requestid interface{} `json:"requestId"`
	Tags interface{} `json:"tags,omitempty"`
}

// FormatToHeader represents the FormatToHeader schema from the OpenAPI specification
type FormatToHeader struct {
}

// CompleteReadSetUploadPartListItem represents the CompleteReadSetUploadPartListItem schema from the OpenAPI specification
type CompleteReadSetUploadPartListItem struct {
	Partsource interface{} `json:"partSource"`
	Checksum interface{} `json:"checksum"`
	Partnumber interface{} `json:"partNumber"`
}

// GetRunResponse represents the GetRunResponse schema from the OpenAPI specification
type GetRunResponse struct {
	Workflowtype interface{} `json:"workflowType,omitempty"`
	Digest interface{} `json:"digest,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Parameters interface{} `json:"parameters,omitempty"`
	Creationtime interface{} `json:"creationTime,omitempty"`
	Startedby interface{} `json:"startedBy,omitempty"`
	Accelerators interface{} `json:"accelerators,omitempty"`
	Runid interface{} `json:"runId,omitempty"`
	Definition interface{} `json:"definition,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Stoptime interface{} `json:"stopTime,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Storagecapacity interface{} `json:"storageCapacity,omitempty"`
	Outputuri interface{} `json:"outputUri,omitempty"`
	Rolearn interface{} `json:"roleArn,omitempty"`
	Statusmessage interface{} `json:"statusMessage,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Loglevel interface{} `json:"logLevel,omitempty"`
	Workflowid interface{} `json:"workflowId,omitempty"`
	Rungroupid interface{} `json:"runGroupId,omitempty"`
	Priority interface{} `json:"priority,omitempty"`
	Resourcedigests interface{} `json:"resourceDigests,omitempty"`
	Starttime interface{} `json:"startTime,omitempty"`
}

// ImportReferenceJobItem represents the ImportReferenceJobItem schema from the OpenAPI specification
type ImportReferenceJobItem struct {
	Referencestoreid interface{} `json:"referenceStoreId"`
	Rolearn interface{} `json:"roleArn"`
	Status interface{} `json:"status"`
	Completiontime interface{} `json:"completionTime,omitempty"`
	Creationtime interface{} `json:"creationTime"`
	Id interface{} `json:"id"`
}

// CreateVariantStoreRequest represents the CreateVariantStoreRequest schema from the OpenAPI specification
type CreateVariantStoreRequest struct {
	Reference interface{} `json:"reference"`
	Sseconfig interface{} `json:"sseConfig,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// GetReferenceStoreResponse represents the GetReferenceStoreResponse schema from the OpenAPI specification
type GetReferenceStoreResponse struct {
	Id interface{} `json:"id"`
	Name interface{} `json:"name,omitempty"`
	Sseconfig interface{} `json:"sseConfig,omitempty"`
	Arn interface{} `json:"arn"`
	Creationtime interface{} `json:"creationTime"`
	Description interface{} `json:"description,omitempty"`
}

// SourceFiles represents the SourceFiles schema from the OpenAPI specification
type SourceFiles struct {
	Source1 interface{} `json:"source1"`
	Source2 interface{} `json:"source2,omitempty"`
}

// CreateMultipartReadSetUploadRequest represents the CreateMultipartReadSetUploadRequest schema from the OpenAPI specification
type CreateMultipartReadSetUploadRequest struct {
	Name interface{} `json:"name"`
	Generatedfrom interface{} `json:"generatedFrom,omitempty"`
	Sourcefiletype interface{} `json:"sourceFileType"`
	Subjectid interface{} `json:"subjectId"`
	Tags interface{} `json:"tags,omitempty"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
	Sampleid interface{} `json:"sampleId"`
	Description interface{} `json:"description,omitempty"`
	Referencearn interface{} `json:"referenceArn"`
}

// ReferenceFiles represents the ReferenceFiles schema from the OpenAPI specification
type ReferenceFiles struct {
	Index interface{} `json:"index,omitempty"`
	Source interface{} `json:"source,omitempty"`
}

// GetReferenceImportJobResponse represents the GetReferenceImportJobResponse schema from the OpenAPI specification
type GetReferenceImportJobResponse struct {
	Status interface{} `json:"status"`
	Statusmessage interface{} `json:"statusMessage,omitempty"`
	Completiontime interface{} `json:"completionTime,omitempty"`
	Creationtime interface{} `json:"creationTime"`
	Id interface{} `json:"id"`
	Referencestoreid interface{} `json:"referenceStoreId"`
	Rolearn interface{} `json:"roleArn"`
	Sources interface{} `json:"sources"`
}

// ListRunGroupsResponse represents the ListRunGroupsResponse schema from the OpenAPI specification
type ListRunGroupsResponse struct {
	Items interface{} `json:"items,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// ListAnnotationImportJobsRequest represents the ListAnnotationImportJobsRequest schema from the OpenAPI specification
type ListAnnotationImportJobsRequest struct {
	Filter interface{} `json:"filter,omitempty"`
	Ids interface{} `json:"ids,omitempty"`
}

// StartRunRequest represents the StartRunRequest schema from the OpenAPI specification
type StartRunRequest struct {
	Workflowtype interface{} `json:"workflowType,omitempty"`
	Loglevel interface{} `json:"logLevel,omitempty"`
	Parameters interface{} `json:"parameters,omitempty"`
	Priority interface{} `json:"priority,omitempty"`
	Rungroupid interface{} `json:"runGroupId,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Requestid interface{} `json:"requestId"`
	Runid interface{} `json:"runId,omitempty"`
	Storagecapacity interface{} `json:"storageCapacity,omitempty"`
	Workflowid interface{} `json:"workflowId,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Outputuri interface{} `json:"outputUri,omitempty"`
	Rolearn interface{} `json:"roleArn"`
}

// GetReferenceImportJobRequest represents the GetReferenceImportJobRequest schema from the OpenAPI specification
type GetReferenceImportJobRequest struct {
}

// GetRunRequest represents the GetRunRequest schema from the OpenAPI specification
type GetRunRequest struct {
}

// UpdateAnnotationStoreRequest represents the UpdateAnnotationStoreRequest schema from the OpenAPI specification
type UpdateAnnotationStoreRequest struct {
	Description interface{} `json:"description,omitempty"`
}

// AnnotationImportItemSource represents the AnnotationImportItemSource schema from the OpenAPI specification
type AnnotationImportItemSource struct {
	Source interface{} `json:"source"`
}

// ExportReadSetFilter represents the ExportReadSetFilter schema from the OpenAPI specification
type ExportReadSetFilter struct {
	Createdafter interface{} `json:"createdAfter,omitempty"`
	Createdbefore interface{} `json:"createdBefore,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// VariantImportItemSource represents the VariantImportItemSource schema from the OpenAPI specification
type VariantImportItemSource struct {
	Source interface{} `json:"source"`
}

// GetSequenceStoreRequest represents the GetSequenceStoreRequest schema from the OpenAPI specification
type GetSequenceStoreRequest struct {
}

// CreateVariantStoreResponse represents the CreateVariantStoreResponse schema from the OpenAPI specification
type CreateVariantStoreResponse struct {
	Id interface{} `json:"id"`
	Name interface{} `json:"name"`
	Reference interface{} `json:"reference,omitempty"`
	Status interface{} `json:"status"`
	Creationtime interface{} `json:"creationTime"`
}

// ListSequenceStoresRequest represents the ListSequenceStoresRequest schema from the OpenAPI specification
type ListSequenceStoresRequest struct {
	Filter interface{} `json:"filter,omitempty"`
}

// UploadReadSetPartResponse represents the UploadReadSetPartResponse schema from the OpenAPI specification
type UploadReadSetPartResponse struct {
	Checksum interface{} `json:"checksum"`
}

// ListReadSetActivationJobsResponse represents the ListReadSetActivationJobsResponse schema from the OpenAPI specification
type ListReadSetActivationJobsResponse struct {
	Activationjobs interface{} `json:"activationJobs,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// StartReferenceImportJobResponse represents the StartReferenceImportJobResponse schema from the OpenAPI specification
type StartReferenceImportJobResponse struct {
	Creationtime interface{} `json:"creationTime"`
	Id interface{} `json:"id"`
	Referencestoreid interface{} `json:"referenceStoreId"`
	Rolearn interface{} `json:"roleArn"`
	Status interface{} `json:"status"`
}

// CancelRunRequest represents the CancelRunRequest schema from the OpenAPI specification
type CancelRunRequest struct {
}

// GetRunTaskResponse represents the GetRunTaskResponse schema from the OpenAPI specification
type GetRunTaskResponse struct {
	Cpus interface{} `json:"cpus,omitempty"`
	Creationtime interface{} `json:"creationTime,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Memory interface{} `json:"memory,omitempty"`
	Statusmessage interface{} `json:"statusMessage,omitempty"`
	Taskid interface{} `json:"taskId,omitempty"`
	Logstream interface{} `json:"logStream,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Starttime interface{} `json:"startTime,omitempty"`
	Gpus interface{} `json:"gpus,omitempty"`
	Stoptime interface{} `json:"stopTime,omitempty"`
}

// ActivateReadSetFilter represents the ActivateReadSetFilter schema from the OpenAPI specification
type ActivateReadSetFilter struct {
	Status interface{} `json:"status,omitempty"`
	Createdafter interface{} `json:"createdAfter,omitempty"`
	Createdbefore interface{} `json:"createdBefore,omitempty"`
}

// GetReadSetImportJobRequest represents the GetReadSetImportJobRequest schema from the OpenAPI specification
type GetReadSetImportJobRequest struct {
}

// ListReadSetActivationJobsRequest represents the ListReadSetActivationJobsRequest schema from the OpenAPI specification
type ListReadSetActivationJobsRequest struct {
	Filter interface{} `json:"filter,omitempty"`
}

// StartRunResponse represents the StartRunResponse schema from the OpenAPI specification
type StartRunResponse struct {
	Arn interface{} `json:"arn,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
}

// StartReferenceImportJobRequest represents the StartReferenceImportJobRequest schema from the OpenAPI specification
type StartReferenceImportJobRequest struct {
	Rolearn interface{} `json:"roleArn"`
	Sources interface{} `json:"sources"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
}

// GetReferenceResponse represents the GetReferenceResponse schema from the OpenAPI specification
type GetReferenceResponse struct {
	Payload interface{} `json:"payload,omitempty"`
}

// GetVariantImportRequest represents the GetVariantImportRequest schema from the OpenAPI specification
type GetVariantImportRequest struct {
}

// AnnotationStoreItem represents the AnnotationStoreItem schema from the OpenAPI specification
type AnnotationStoreItem struct {
	Sseconfig interface{} `json:"sseConfig"`
	Updatetime interface{} `json:"updateTime"`
	Creationtime interface{} `json:"creationTime"`
	Description interface{} `json:"description"`
	Name interface{} `json:"name"`
	Statusmessage interface{} `json:"statusMessage"`
	Storearn interface{} `json:"storeArn"`
	Reference interface{} `json:"reference"`
	Storeformat interface{} `json:"storeFormat"`
	Id interface{} `json:"id"`
	Status interface{} `json:"status"`
	Storesizebytes interface{} `json:"storeSizeBytes"`
}

// MultipartReadSetUploadListItem represents the MultipartReadSetUploadListItem schema from the OpenAPI specification
type MultipartReadSetUploadListItem struct {
	Sourcefiletype interface{} `json:"sourceFileType"`
	Description interface{} `json:"description,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Referencearn interface{} `json:"referenceArn"`
	Subjectid interface{} `json:"subjectId"`
	Tags interface{} `json:"tags,omitempty"`
	Uploadid interface{} `json:"uploadId"`
	Creationtime interface{} `json:"creationTime"`
	Generatedfrom interface{} `json:"generatedFrom"`
	Sampleid interface{} `json:"sampleId"`
	Sequencestoreid interface{} `json:"sequenceStoreId"`
}

// ReadSetBatchError represents the ReadSetBatchError schema from the OpenAPI specification
type ReadSetBatchError struct {
	Code interface{} `json:"code"`
	Id interface{} `json:"id"`
	Message interface{} `json:"message"`
}

// ListRunsResponse represents the ListRunsResponse schema from the OpenAPI specification
type ListRunsResponse struct {
	Items interface{} `json:"items,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// DeleteRunGroupRequest represents the DeleteRunGroupRequest schema from the OpenAPI specification
type DeleteRunGroupRequest struct {
}

// UpdateWorkflowRequest represents the UpdateWorkflowRequest schema from the OpenAPI specification
type UpdateWorkflowRequest struct {
	Description interface{} `json:"description,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// GetSequenceStoreResponse represents the GetSequenceStoreResponse schema from the OpenAPI specification
type GetSequenceStoreResponse struct {
	Fallbacklocation interface{} `json:"fallbackLocation,omitempty"`
	Id interface{} `json:"id"`
	Name interface{} `json:"name,omitempty"`
	Sseconfig interface{} `json:"sseConfig,omitempty"`
	Arn interface{} `json:"arn"`
	Creationtime interface{} `json:"creationTime"`
	Description interface{} `json:"description,omitempty"`
}

// SseConfig represents the SseConfig schema from the OpenAPI specification
type SseConfig struct {
	Keyarn interface{} `json:"keyArn,omitempty"`
	TypeField interface{} `json:"type"`
}

// ExportReadSet represents the ExportReadSet schema from the OpenAPI specification
type ExportReadSet struct {
	Readsetid interface{} `json:"readSetId"`
}

// WorkflowParameter represents the WorkflowParameter schema from the OpenAPI specification
type WorkflowParameter struct {
	Optional interface{} `json:"optional,omitempty"`
	Description interface{} `json:"description,omitempty"`
}

// WorkflowParameterTemplate represents the WorkflowParameterTemplate schema from the OpenAPI specification
type WorkflowParameterTemplate struct {
}

// ListAnnotationStoresRequest represents the ListAnnotationStoresRequest schema from the OpenAPI specification
type ListAnnotationStoresRequest struct {
	Ids interface{} `json:"ids,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

// VariantStoreItem represents the VariantStoreItem schema from the OpenAPI specification
type VariantStoreItem struct {
	Creationtime interface{} `json:"creationTime"`
	Reference interface{} `json:"reference"`
	Storearn interface{} `json:"storeArn"`
	Updatetime interface{} `json:"updateTime"`
	Description interface{} `json:"description"`
	Sseconfig interface{} `json:"sseConfig"`
	Status interface{} `json:"status"`
	Storesizebytes interface{} `json:"storeSizeBytes"`
	Id interface{} `json:"id"`
	Name interface{} `json:"name"`
	Statusmessage interface{} `json:"statusMessage"`
}

// TaskListItem represents the TaskListItem schema from the OpenAPI specification
type TaskListItem struct {
	Creationtime interface{} `json:"creationTime,omitempty"`
	Memory interface{} `json:"memory,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Starttime interface{} `json:"startTime,omitempty"`
	Cpus interface{} `json:"cpus,omitempty"`
	Gpus interface{} `json:"gpus,omitempty"`
	Stoptime interface{} `json:"stopTime,omitempty"`
	Taskid interface{} `json:"taskId,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// GetReadSetActivationJobResponse represents the GetReadSetActivationJobResponse schema from the OpenAPI specification
type GetReadSetActivationJobResponse struct {
	Sources interface{} `json:"sources,omitempty"`
	Status interface{} `json:"status"`
	Statusmessage interface{} `json:"statusMessage,omitempty"`
	Completiontime interface{} `json:"completionTime,omitempty"`
	Creationtime interface{} `json:"creationTime"`
	Id interface{} `json:"id"`
	Sequencestoreid interface{} `json:"sequenceStoreId"`
}

// GetRunTaskRequest represents the GetRunTaskRequest schema from the OpenAPI specification
type GetRunTaskRequest struct {
}

// BatchDeleteReadSetRequest represents the BatchDeleteReadSetRequest schema from the OpenAPI specification
type BatchDeleteReadSetRequest struct {
	Ids interface{} `json:"ids"`
}

// ListRunGroupsRequest represents the ListRunGroupsRequest schema from the OpenAPI specification
type ListRunGroupsRequest struct {
}

// ReferenceFilter represents the ReferenceFilter schema from the OpenAPI specification
type ReferenceFilter struct {
	Createdafter interface{} `json:"createdAfter,omitempty"`
	Createdbefore interface{} `json:"createdBefore,omitempty"`
	Md5 interface{} `json:"md5,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// ImportReferenceFilter represents the ImportReferenceFilter schema from the OpenAPI specification
type ImportReferenceFilter struct {
	Createdafter interface{} `json:"createdAfter,omitempty"`
	Createdbefore interface{} `json:"createdBefore,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// ImportReadSetSourceItem represents the ImportReadSetSourceItem schema from the OpenAPI specification
type ImportReadSetSourceItem struct {
	Name interface{} `json:"name,omitempty"`
	Sampleid interface{} `json:"sampleId"`
	Sourcefiles interface{} `json:"sourceFiles"`
	Statusmessage interface{} `json:"statusMessage,omitempty"`
	Sourcefiletype interface{} `json:"sourceFileType"`
	Subjectid interface{} `json:"subjectId"`
	Tags interface{} `json:"tags,omitempty"`
	Referencearn interface{} `json:"referenceArn,omitempty"`
	Status interface{} `json:"status"`
	Description interface{} `json:"description,omitempty"`
	Generatedfrom interface{} `json:"generatedFrom,omitempty"`
}

// StartReadSetImportJobSourceItem represents the StartReadSetImportJobSourceItem schema from the OpenAPI specification
type StartReadSetImportJobSourceItem struct {
	Description interface{} `json:"description,omitempty"`
	Sourcefiletype interface{} `json:"sourceFileType"`
	Subjectid interface{} `json:"subjectId"`
	Generatedfrom interface{} `json:"generatedFrom,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Sourcefiles interface{} `json:"sourceFiles"`
	Referencearn interface{} `json:"referenceArn"`
	Sampleid interface{} `json:"sampleId"`
	Tags interface{} `json:"tags,omitempty"`
}

// ListVariantStoresResponse represents the ListVariantStoresResponse schema from the OpenAPI specification
type ListVariantStoresResponse struct {
	Variantstores interface{} `json:"variantStores,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// StartAnnotationImportResponse represents the StartAnnotationImportResponse schema from the OpenAPI specification
type StartAnnotationImportResponse struct {
	Jobid interface{} `json:"jobId"`
}

// TsvStoreOptions represents the TsvStoreOptions schema from the OpenAPI specification
type TsvStoreOptions struct {
	Annotationtype interface{} `json:"annotationType,omitempty"`
	Formattoheader interface{} `json:"formatToHeader,omitempty"`
	Schema interface{} `json:"schema,omitempty"`
}

// UpdateVariantStoreResponse represents the UpdateVariantStoreResponse schema from the OpenAPI specification
type UpdateVariantStoreResponse struct {
	Description interface{} `json:"description"`
	Id interface{} `json:"id"`
	Name interface{} `json:"name"`
	Reference interface{} `json:"reference"`
	Status interface{} `json:"status"`
	Updatetime interface{} `json:"updateTime"`
	Creationtime interface{} `json:"creationTime"`
}

// UploadReadSetPartRequest represents the UploadReadSetPartRequest schema from the OpenAPI specification
type UploadReadSetPartRequest struct {
	Payload interface{} `json:"payload"`
}

// DeleteWorkflowRequest represents the DeleteWorkflowRequest schema from the OpenAPI specification
type DeleteWorkflowRequest struct {
}

// ListAnnotationImportJobsResponse represents the ListAnnotationImportJobsResponse schema from the OpenAPI specification
type ListAnnotationImportJobsResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Annotationimportjobs interface{} `json:"annotationImportJobs,omitempty"`
}

// FileInformation represents the FileInformation schema from the OpenAPI specification
type FileInformation struct {
	Partsize interface{} `json:"partSize,omitempty"`
	Totalparts interface{} `json:"totalParts,omitempty"`
	Contentlength interface{} `json:"contentLength,omitempty"`
}

// SequenceInformation represents the SequenceInformation schema from the OpenAPI specification
type SequenceInformation struct {
	Generatedfrom interface{} `json:"generatedFrom,omitempty"`
	Totalbasecount interface{} `json:"totalBaseCount,omitempty"`
	Totalreadcount interface{} `json:"totalReadCount,omitempty"`
	Alignment interface{} `json:"alignment,omitempty"`
}

// DeleteReferenceResponse represents the DeleteReferenceResponse schema from the OpenAPI specification
type DeleteReferenceResponse struct {
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
}

// ListReadSetImportJobsRequest represents the ListReadSetImportJobsRequest schema from the OpenAPI specification
type ListReadSetImportJobsRequest struct {
	Filter interface{} `json:"filter,omitempty"`
}

// GetReadSetActivationJobRequest represents the GetReadSetActivationJobRequest schema from the OpenAPI specification
type GetReadSetActivationJobRequest struct {
}

// ReadSetUploadPartListFilter represents the ReadSetUploadPartListFilter schema from the OpenAPI specification
type ReadSetUploadPartListFilter struct {
	Createdafter interface{} `json:"createdAfter,omitempty"`
	Createdbefore interface{} `json:"createdBefore,omitempty"`
}

// DeleteAnnotationStoreResponse represents the DeleteAnnotationStoreResponse schema from the OpenAPI specification
type DeleteAnnotationStoreResponse struct {
	Status interface{} `json:"status"`
}

// StartVariantImportRequest represents the StartVariantImportRequest schema from the OpenAPI specification
type StartVariantImportRequest struct {
	Items interface{} `json:"items"`
	Rolearn interface{} `json:"roleArn"`
	Runleftnormalization interface{} `json:"runLeftNormalization,omitempty"`
	Annotationfields interface{} `json:"annotationFields,omitempty"`
	Destinationname interface{} `json:"destinationName"`
}

// StartReadSetImportJobResponse represents the StartReadSetImportJobResponse schema from the OpenAPI specification
type StartReadSetImportJobResponse struct {
	Creationtime interface{} `json:"creationTime"`
	Id interface{} `json:"id"`
	Rolearn interface{} `json:"roleArn"`
	Sequencestoreid interface{} `json:"sequenceStoreId"`
	Status interface{} `json:"status"`
}

// RunListItem represents the RunListItem schema from the OpenAPI specification
type RunListItem struct {
	Arn interface{} `json:"arn,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Starttime interface{} `json:"startTime,omitempty"`
	Stoptime interface{} `json:"stopTime,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Storagecapacity interface{} `json:"storageCapacity,omitempty"`
	Creationtime interface{} `json:"creationTime,omitempty"`
	Workflowid interface{} `json:"workflowId,omitempty"`
	Priority interface{} `json:"priority,omitempty"`
}

// DeleteReferenceStoreRequest represents the DeleteReferenceStoreRequest schema from the OpenAPI specification
type DeleteReferenceStoreRequest struct {
}

// DeleteRunRequest represents the DeleteRunRequest schema from the OpenAPI specification
type DeleteRunRequest struct {
}

// ReferenceListItem represents the ReferenceListItem schema from the OpenAPI specification
type ReferenceListItem struct {
	Status interface{} `json:"status,omitempty"`
	Creationtime interface{} `json:"creationTime"`
	Id interface{} `json:"id"`
	Description interface{} `json:"description,omitempty"`
	Referencestoreid interface{} `json:"referenceStoreId"`
	Updatetime interface{} `json:"updateTime"`
	Arn interface{} `json:"arn"`
	Md5 interface{} `json:"md5"`
	Name interface{} `json:"name,omitempty"`
}

// GetAnnotationImportRequest represents the GetAnnotationImportRequest schema from the OpenAPI specification
type GetAnnotationImportRequest struct {
}

// WorkflowListItem represents the WorkflowListItem schema from the OpenAPI specification
type WorkflowListItem struct {
	Name interface{} `json:"name,omitempty"`
	Status interface{} `json:"status,omitempty"`
	TypeField interface{} `json:"type,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Creationtime interface{} `json:"creationTime,omitempty"`
	Digest interface{} `json:"digest,omitempty"`
	Id interface{} `json:"id,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
}

// ListReadSetUploadPartsResponse represents the ListReadSetUploadPartsResponse schema from the OpenAPI specification
type ListReadSetUploadPartsResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Parts interface{} `json:"parts,omitempty"`
}

// RunGroupListItem represents the RunGroupListItem schema from the OpenAPI specification
type RunGroupListItem struct {
	Id interface{} `json:"id,omitempty"`
	Maxcpus interface{} `json:"maxCpus,omitempty"`
	Maxduration interface{} `json:"maxDuration,omitempty"`
	Maxgpus interface{} `json:"maxGpus,omitempty"`
	Maxruns interface{} `json:"maxRuns,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Arn interface{} `json:"arn,omitempty"`
	Creationtime interface{} `json:"creationTime,omitempty"`
}

// GetReadSetImportJobResponse represents the GetReadSetImportJobResponse schema from the OpenAPI specification
type GetReadSetImportJobResponse struct {
	Creationtime interface{} `json:"creationTime"`
	Id interface{} `json:"id"`
	Rolearn interface{} `json:"roleArn"`
	Sequencestoreid interface{} `json:"sequenceStoreId"`
	Sources interface{} `json:"sources"`
	Status interface{} `json:"status"`
	Statusmessage interface{} `json:"statusMessage,omitempty"`
	Completiontime interface{} `json:"completionTime,omitempty"`
}

// CreateWorkflowRequest represents the CreateWorkflowRequest schema from the OpenAPI specification
type CreateWorkflowRequest struct {
	Accelerators interface{} `json:"accelerators,omitempty"`
	Definitionzip interface{} `json:"definitionZip,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Engine interface{} `json:"engine,omitempty"`
	Definitionuri interface{} `json:"definitionUri,omitempty"`
	Requestid interface{} `json:"requestId"`
	Tags interface{} `json:"tags,omitempty"`
	Main interface{} `json:"main,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Parametertemplate interface{} `json:"parameterTemplate,omitempty"`
	Storagecapacity interface{} `json:"storageCapacity,omitempty"`
}

// SequenceStoreFilter represents the SequenceStoreFilter schema from the OpenAPI specification
type SequenceStoreFilter struct {
	Name interface{} `json:"name,omitempty"`
	Createdafter interface{} `json:"createdAfter,omitempty"`
	Createdbefore interface{} `json:"createdBefore,omitempty"`
}

// CreateSequenceStoreResponse represents the CreateSequenceStoreResponse schema from the OpenAPI specification
type CreateSequenceStoreResponse struct {
	Id interface{} `json:"id"`
	Name interface{} `json:"name,omitempty"`
	Sseconfig interface{} `json:"sseConfig,omitempty"`
	Arn interface{} `json:"arn"`
	Creationtime interface{} `json:"creationTime"`
	Description interface{} `json:"description,omitempty"`
	Fallbacklocation interface{} `json:"fallbackLocation,omitempty"`
}

// DeleteVariantStoreResponse represents the DeleteVariantStoreResponse schema from the OpenAPI specification
type DeleteVariantStoreResponse struct {
	Status interface{} `json:"status"`
}

// SequenceStoreDetail represents the SequenceStoreDetail schema from the OpenAPI specification
type SequenceStoreDetail struct {
	Description interface{} `json:"description,omitempty"`
	Fallbacklocation interface{} `json:"fallbackLocation,omitempty"`
	Id interface{} `json:"id"`
	Name interface{} `json:"name,omitempty"`
	Sseconfig interface{} `json:"sseConfig,omitempty"`
	Arn interface{} `json:"arn"`
	Creationtime interface{} `json:"creationTime"`
}

// GetReadSetMetadataResponse represents the GetReadSetMetadataResponse schema from the OpenAPI specification
type GetReadSetMetadataResponse struct {
	Files interface{} `json:"files,omitempty"`
	Status interface{} `json:"status"`
	Arn interface{} `json:"arn"`
	Creationtype interface{} `json:"creationType,omitempty"`
	Filetype interface{} `json:"fileType"`
	Sequencestoreid interface{} `json:"sequenceStoreId"`
	Creationtime interface{} `json:"creationTime"`
	Id interface{} `json:"id"`
	Referencearn interface{} `json:"referenceArn,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Sampleid interface{} `json:"sampleId,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Statusmessage interface{} `json:"statusMessage,omitempty"`
	Subjectid interface{} `json:"subjectId,omitempty"`
	Sequenceinformation interface{} `json:"sequenceInformation,omitempty"`
}

// ListAnnotationStoresFilter represents the ListAnnotationStoresFilter schema from the OpenAPI specification
type ListAnnotationStoresFilter struct {
	Status interface{} `json:"status,omitempty"`
}

// DeleteReferenceRequest represents the DeleteReferenceRequest schema from the OpenAPI specification
type DeleteReferenceRequest struct {
}

// SchemaItem represents the SchemaItem schema from the OpenAPI specification
type SchemaItem struct {
}

// CancelVariantImportResponse represents the CancelVariantImportResponse schema from the OpenAPI specification
type CancelVariantImportResponse struct {
}

// ReadSetUploadPartListItem represents the ReadSetUploadPartListItem schema from the OpenAPI specification
type ReadSetUploadPartListItem struct {
	Partsize interface{} `json:"partSize"`
	Partsource interface{} `json:"partSource"`
	Checksum interface{} `json:"checksum"`
	Creationtime interface{} `json:"creationTime,omitempty"`
	Lastupdatedtime interface{} `json:"lastUpdatedTime,omitempty"`
	Partnumber interface{} `json:"partNumber"`
}

// DeleteSequenceStoreResponse represents the DeleteSequenceStoreResponse schema from the OpenAPI specification
type DeleteSequenceStoreResponse struct {
}

// ListRunTasksResponse represents the ListRunTasksResponse schema from the OpenAPI specification
type ListRunTasksResponse struct {
	Items interface{} `json:"items,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// ListRunTasksRequest represents the ListRunTasksRequest schema from the OpenAPI specification
type ListRunTasksRequest struct {
}

// ListReferenceImportJobsRequest represents the ListReferenceImportJobsRequest schema from the OpenAPI specification
type ListReferenceImportJobsRequest struct {
	Filter interface{} `json:"filter,omitempty"`
}

// ReadOptions represents the ReadOptions schema from the OpenAPI specification
type ReadOptions struct {
	Quote interface{} `json:"quote,omitempty"`
	Quoteall interface{} `json:"quoteAll,omitempty"`
	Encoding interface{} `json:"encoding,omitempty"`
	Escape interface{} `json:"escape,omitempty"`
	Escapequotes interface{} `json:"escapeQuotes,omitempty"`
	Sep interface{} `json:"sep,omitempty"`
	Header interface{} `json:"header,omitempty"`
	Linesep interface{} `json:"lineSep,omitempty"`
	Comment interface{} `json:"comment,omitempty"`
}

// ListMultipartReadSetUploadsRequest represents the ListMultipartReadSetUploadsRequest schema from the OpenAPI specification
type ListMultipartReadSetUploadsRequest struct {
}

// ListAnnotationStoresResponse represents the ListAnnotationStoresResponse schema from the OpenAPI specification
type ListAnnotationStoresResponse struct {
	Annotationstores interface{} `json:"annotationStores,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// VcfOptions represents the VcfOptions schema from the OpenAPI specification
type VcfOptions struct {
	Ignorefilterfield interface{} `json:"ignoreFilterField,omitempty"`
	Ignorequalfield interface{} `json:"ignoreQualField,omitempty"`
}

// GetRunGroupRequest represents the GetRunGroupRequest schema from the OpenAPI specification
type GetRunGroupRequest struct {
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Tags interface{} `json:"tags"`
}

// GetVariantStoreResponse represents the GetVariantStoreResponse schema from the OpenAPI specification
type GetVariantStoreResponse struct {
	Creationtime interface{} `json:"creationTime"`
	Description interface{} `json:"description"`
	Sseconfig interface{} `json:"sseConfig"`
	Statusmessage interface{} `json:"statusMessage"`
	Tags interface{} `json:"tags"`
	Id interface{} `json:"id"`
	Name interface{} `json:"name"`
	Storearn interface{} `json:"storeArn"`
	Status interface{} `json:"status"`
	Updatetime interface{} `json:"updateTime"`
	Reference interface{} `json:"reference"`
	Storesizebytes interface{} `json:"storeSizeBytes"`
}

// ListReadSetImportJobsResponse represents the ListReadSetImportJobsResponse schema from the OpenAPI specification
type ListReadSetImportJobsResponse struct {
	Importjobs interface{} `json:"importJobs,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// WorkflowMetadata represents the WorkflowMetadata schema from the OpenAPI specification
type WorkflowMetadata struct {
}

// ReadSetFiles represents the ReadSetFiles schema from the OpenAPI specification
type ReadSetFiles struct {
	Index interface{} `json:"index,omitempty"`
	Source1 interface{} `json:"source1,omitempty"`
	Source2 interface{} `json:"source2,omitempty"`
}

// DeleteReferenceStoreResponse represents the DeleteReferenceStoreResponse schema from the OpenAPI specification
type DeleteReferenceStoreResponse struct {
}

// UpdateAnnotationStoreResponse represents the UpdateAnnotationStoreResponse schema from the OpenAPI specification
type UpdateAnnotationStoreResponse struct {
	Description interface{} `json:"description"`
	Name interface{} `json:"name"`
	Storeformat interface{} `json:"storeFormat,omitempty"`
	Id interface{} `json:"id"`
	Reference interface{} `json:"reference"`
	Status interface{} `json:"status"`
	Storeoptions interface{} `json:"storeOptions,omitempty"`
	Updatetime interface{} `json:"updateTime"`
	Creationtime interface{} `json:"creationTime"`
}

// StartReadSetExportJobRequest represents the StartReadSetExportJobRequest schema from the OpenAPI specification
type StartReadSetExportJobRequest struct {
	Destination interface{} `json:"destination"`
	Rolearn interface{} `json:"roleArn"`
	Sources interface{} `json:"sources"`
	Clienttoken interface{} `json:"clientToken,omitempty"`
}

// ListMultipartReadSetUploadsResponse represents the ListMultipartReadSetUploadsResponse schema from the OpenAPI specification
type ListMultipartReadSetUploadsResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Uploads interface{} `json:"uploads,omitempty"`
}

// GetReadSetMetadataRequest represents the GetReadSetMetadataRequest schema from the OpenAPI specification
type GetReadSetMetadataRequest struct {
}

// ListRunsRequest represents the ListRunsRequest schema from the OpenAPI specification
type ListRunsRequest struct {
}

// ReferenceStoreFilter represents the ReferenceStoreFilter schema from the OpenAPI specification
type ReferenceStoreFilter struct {
	Name interface{} `json:"name,omitempty"`
	Createdafter interface{} `json:"createdAfter,omitempty"`
	Createdbefore interface{} `json:"createdBefore,omitempty"`
}
