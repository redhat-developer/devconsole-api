package v1alpha1

// GetAnnotationVcsUri retrieves the annotation vcs-uri.
func (c *Component) GetAnnotationVcsUri() string {
	return c.Annotations["app.openshift.io/vcs-uri"]
}

// GetAnnotationVcsRef retrieves the annotation vcs-red.
func (c *Component) GetAnnotationVcsRef() string {
	return c.Annotations["app.openshift.io/vcs-ref"]
}
