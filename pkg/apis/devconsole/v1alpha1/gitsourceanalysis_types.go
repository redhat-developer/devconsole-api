package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GitSourceAnalysisSpec defines the desired state of GitSourceAnalysis
type GitSourceAnalysisSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file

	// GitSourceRef refers to the GitSource to be analyzed
	GitSourceRef GitSourceRef `json:"gitSourceRef"`
}

// GitSourceRef refers to the GitSource to be analyzed
type GitSourceRef struct {
	// Name is the name of the GitSource within the same namespace that contains all necessary information of the git repo
	Name string `json:"name"`
}

// GitSourceAnalysisStatus defines the observed state of GitSourceAnalysis
type GitSourceAnalysisStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file

	// Analyzed says if the GitSource analysis is done or not
	Analyzed bool `json:"analyzed"`

	// BuildEnvStatistics holds information about detected languages and build types in the GitSource. Optional
	BuildEnvStatistics BuildEnvStats `json:"buildEnvStatistics,omitempty"`

	// Error contains an error message in case the build environment detection fails. Optional
	Error string `json:"error,omitempty"`

	// Reason represents the reason why the GitSource analysis (build type detection) failed. Optional
	Reason AnalysisFailureReason `json:"reason,omitempty"`
}

// AnalysisFailureReason represents the reason why the GitSource analysis (build type detection) failed
type AnalysisFailureReason string

const (
	// NotSupportedType represents the failure reason when no appropriate git implementation was found
	// for the given combination of repository and secret
	NotSupportedType AnalysisFailureReason = "NotSupportedType"
	// DetectionFailed represents the failure reason when the actual detection logic failed
	DetectionFailed AnalysisFailureReason = "DetectionFailed"
	// AnalysisInternalFailure represents a failure reason caused by any internal failure
	AnalysisInternalFailure AnalysisFailureReason = "InternalFailure"
)

// BuildEnvStatistics holds information about detected languages and build types in the GitSource
type BuildEnvStats struct {
	// SortedLanguages contains sorted languages detected in the GitSource where the first one is with the most used
	SortedLanguages []string `json:"sortedLanguages,omitempty"`

	// DetectedBuildTypes contains list of detected build types in the GitSource
	DetectedBuildTypes []DetectedBuildType `json:"detectedBuildTypes,omitempty"`
}

// DetectedBuildType holds information of the build type detected in a GitSource
type DetectedBuildType struct {
	// Language is a programing language the build type if used for
	Language string `json:"language,omitempty"`

	// Name is a name of the build type
	Name string `json:"name,omitempty"`

	// DetectedFiles contains a list of files detected in the GitSource that are used by the build type
	DetectedFiles []string `json:"detectedFiles,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GitSourceAnalysis is the Schema for the gitsourceanalyses API
// +k8s:openapi-gen=true
type GitSourceAnalysis struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GitSourceAnalysisSpec   `json:"spec,omitempty"`
	Status GitSourceAnalysisStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GitSourceAnalysisList contains a list of GitSourceAnalysis
type GitSourceAnalysisList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GitSourceAnalysis `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GitSourceAnalysis{}, &GitSourceAnalysisList{})
}
