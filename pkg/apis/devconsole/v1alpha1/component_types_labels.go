package v1alpha1

// GetLabelName retrieves the label name. If not present,
// infer the name by using the component name value.
func (c *Component) GetLabelName() string {
	name := c.Labels["app.kubernetes.io/name"]
	if name == "" {
		name = c.Name
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
		instance = c.GetLabelName()
	}
	return instance
}

// GetLabelVersion retrieves label version.
func (c *Component) GetLabelVersion() string {
	return c.Labels["app.kubernetes.io/version"]
}
