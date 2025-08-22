package main

import (
	"github.com/amazon-omics/mcp-server/config"
	"github.com/amazon-omics/mcp-server/models"
	tools_referencestore "github.com/amazon-omics/mcp-server/tools/referencestore"
	tools_sequencestore "github.com/amazon-omics/mcp-server/tools/sequencestore"
	tools_variantstore "github.com/amazon-omics/mcp-server/tools/variantstore"
	tools_import "github.com/amazon-omics/mcp-server/tools/import"
	tools_run "github.com/amazon-omics/mcp-server/tools/run"
	tools_variantstores "github.com/amazon-omics/mcp-server/tools/variantstores"
	tools_sequencestores "github.com/amazon-omics/mcp-server/tools/sequencestores"
	tools_annotationstore "github.com/amazon-omics/mcp-server/tools/annotationstore"
	tools_tags "github.com/amazon-omics/mcp-server/tools/tags"
	tools_workflow "github.com/amazon-omics/mcp-server/tools/workflow"
	tools_annotationstores "github.com/amazon-omics/mcp-server/tools/annotationstores"
	tools_referencestores "github.com/amazon-omics/mcp-server/tools/referencestores"
	tools_rungroup "github.com/amazon-omics/mcp-server/tools/rungroup"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_referencestore.CreateCreatereferencestoreTool(cfg),
		tools_sequencestore.CreateListreadsetuploadpartsTool(cfg),
		tools_variantstore.CreateDeletevariantstoreTool(cfg),
		tools_variantstore.CreateGetvariantstoreTool(cfg),
		tools_variantstore.CreateUpdatevariantstoreTool(cfg),
		tools_import.CreateListvariantimportjobsTool(cfg),
		tools_run.CreateListruntasksTool(cfg),
		tools_run.CreateGetruntaskTool(cfg),
		tools_referencestore.CreateGetreferenceTool(cfg),
		tools_referencestore.CreateListreferenceimportjobsTool(cfg),
		tools_run.CreateListrunsTool(cfg),
		tools_run.CreateStartrunTool(cfg),
		tools_variantstores.CreateListvariantstoresTool(cfg),
		tools_referencestore.CreateListreferencesTool(cfg),
		tools_sequencestores.CreateListsequencestoresTool(cfg),
		tools_sequencestore.CreateDeletesequencestoreTool(cfg),
		tools_sequencestore.CreateGetsequencestoreTool(cfg),
		tools_sequencestore.CreateStartreadsetimportjobTool(cfg),
		tools_referencestore.CreateDeletereferenceTool(cfg),
		tools_import.CreateListannotationimportjobsTool(cfg),
		tools_referencestore.CreateGetreferenceimportjobTool(cfg),
		tools_run.CreateDeleterunTool(cfg),
		tools_run.CreateGetrunTool(cfg),
		tools_sequencestore.CreateGetreadsetactivationjobTool(cfg),
		tools_variantstore.CreateCreatevariantstoreTool(cfg),
		tools_sequencestore.CreateStartreadsetactivationjobTool(cfg),
		tools_sequencestore.CreateUploadreadsetpartTool(cfg),
		tools_import.CreateStartvariantimportjobTool(cfg),
		tools_run.CreateCancelrunTool(cfg),
		tools_sequencestore.CreateListreadsetimportjobsTool(cfg),
		tools_sequencestore.CreateBatchdeletereadsetTool(cfg),
		tools_sequencestore.CreateGetreadsetexportjobTool(cfg),
		tools_sequencestore.CreateGetreadsetmetadataTool(cfg),
		tools_annotationstore.CreateDeleteannotationstoreTool(cfg),
		tools_annotationstore.CreateGetannotationstoreTool(cfg),
		tools_annotationstore.CreateUpdateannotationstoreTool(cfg),
		tools_import.CreateCancelannotationimportjobTool(cfg),
		tools_import.CreateGetannotationimportjobTool(cfg),
		tools_tags.CreateListtagsforresourceTool(cfg),
		tools_tags.CreateTagresourceTool(cfg),
		tools_referencestore.CreateGetreferencemetadataTool(cfg),
		tools_workflow.CreateDeleteworkflowTool(cfg),
		tools_workflow.CreateGetworkflowTool(cfg),
		tools_workflow.CreateUpdateworkflowTool(cfg),
		tools_workflow.CreateListworkflowsTool(cfg),
		tools_workflow.CreateCreateworkflowTool(cfg),
		tools_sequencestore.CreateGetreadsetimportjobTool(cfg),
		tools_tags.CreateUntagresourceTool(cfg),
		tools_sequencestore.CreateStartreadsetexportjobTool(cfg),
		tools_sequencestore.CreateCompletemultipartreadsetuploadTool(cfg),
		tools_annotationstores.CreateListannotationstoresTool(cfg),
		tools_referencestore.CreateGetreferencestoreTool(cfg),
		tools_referencestore.CreateDeletereferencestoreTool(cfg),
		tools_sequencestore.CreateListreadsetexportjobsTool(cfg),
		tools_sequencestore.CreateListreadsetsTool(cfg),
		tools_sequencestore.CreateAbortmultipartreadsetuploadTool(cfg),
		tools_referencestores.CreateListreferencestoresTool(cfg),
		tools_rungroup.CreateCreaterungroupTool(cfg),
		tools_rungroup.CreateListrungroupsTool(cfg),
		tools_sequencestore.CreateCreatemultipartreadsetuploadTool(cfg),
		tools_rungroup.CreateDeleterungroupTool(cfg),
		tools_rungroup.CreateGetrungroupTool(cfg),
		tools_rungroup.CreateUpdaterungroupTool(cfg),
		tools_sequencestore.CreateGetreadsetTool(cfg),
		tools_import.CreateCancelvariantimportjobTool(cfg),
		tools_import.CreateGetvariantimportjobTool(cfg),
		tools_sequencestore.CreateCreatesequencestoreTool(cfg),
		tools_sequencestore.CreateListmultipartreadsetuploadsTool(cfg),
		tools_sequencestore.CreateListreadsetactivationjobsTool(cfg),
		tools_referencestore.CreateStartreferenceimportjobTool(cfg),
		tools_annotationstore.CreateCreateannotationstoreTool(cfg),
		tools_import.CreateStartannotationimportjobTool(cfg),
	}
}
