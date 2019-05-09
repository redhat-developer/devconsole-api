package v1alpha1

// GetLabelName retrieves the label name. If not present,
// infer the name by using the component spec buildtype value.
func (c *Component) GetLabelName() string {
	name := c.Labels["app.kubernetes.io/name"]
	if name == "" {
		name = c.Spec.BuildType
	}
	return name
}

// GetLabelComponent retrieves the label component.
func (c *Component) GetLabelComponent() string {
	return c.Labels["app.kubernetes.io/component"]
}

// GetLabelPartOf retrieves the label part-of.
func (c *Component) GetLabelPartOf() string {
	return c.Labels["app.kubernetes.io/part-of"]
}

// GetLabelInstance retrieves the label instance. If not present,
// infer the instance by calling GetName().
func (c *Component) GetLabelInstance() string {
	instance := c.Labels["app.kubernetes.io/instance"]
	if instance == "" {
		// Set instance to the component's name.
		// This is used in secondary resource (bc, dc, build) as label selector for a given component.
		instance = c.Name
	}
	return instance
}

// GetLabelVersion retrieves label version.
func (c *Component) GetLabelVersion() string {
	return c.Labels["app.kubernetes.io/version"]
}

// GetLabelDeploymentConfig retrieves the label deploymentconfig. This would be
// same as value of label instance.
func (c *Component) GetLabelDeploymentConfig() string {
	return c.GetLabelInstance()
}
