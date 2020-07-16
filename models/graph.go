package models

// Graph struct is containing all the graph values that will be create on excel sheet
type Graph struct {
	Type     string    `json:"type"`
	Series   []*Series `json:"series"`
	Format   *Format   `json:"format,omitempty"`
	Legend   *Legend   `json:"legend"`
	Title    *Title    `json:"title,omitempty"`
	Plotarea *Plotarea `json:"plotarea,omitempty"`
}

// Series struct is containing all the series of displayed graph
type Series struct {
	Name       string `json:"name"`
	Categories string `json:"categories,omitempty"`
	Values     string `json:"values"`
}

// Format struct is containing all the graph format
type Format struct {
	XScale          float32 `json:"x_scale,omitempty"`
	YScale          float32 `json:"y_scale,omitempty"`
	XOffset         float32 `json:"x_offset,omitempty"`
	YOffset         float32 `json:"y_offset,omitempty"`
	PrintObj        bool    `json:"print_obj,omitempty"`
	LockAspectRatio bool    `json:"lock_aspect_ratio,omitempty"`
	Locked          bool    `json:"locked,omitempty"`
}

// Legend struct is containing graph legend
type Legend struct {
	Position      string `json:"position,omitempty"`
	ShowLegendKey bool   `json:"show_legend_key,omitempty"`
}

// Title struct is containing graph title
type Title struct {
	Name string `json:"name,omitempty"`
}

// Plotarea struct is containing all options for graph
type Plotarea struct {
	ShowBubbleSize  bool `json:"show_bubble_size,omitempty"`
	ShowCatName     bool `json:"show_cat_name,omitempty"`
	ShowLeaderLines bool `json:"show_leader_lines,omitempty"`
	ShowPercent     bool `json:"show_percent,omitempty"`
	ShowSeriesName  bool `json:"show_series_name,omitempty"`
	ShowVal         bool `json:"show_val,omitempty"`
}
