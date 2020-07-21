package models

// Chart struct is containing all the chart values that will be create on excel sheet
type Chart struct {
	Type      string     `json:"type"`
	Series    []*Series  `json:"series"`
	Format    *Format    `json:"format,omitempty"`
	Legend    *Legend    `json:"legend"`
	Title     *Title     `json:"title,omitempty"`
	PlotArea  *PlotArea  `json:"plot_area,omitempty"`
	Dimension *Dimension `json:"dimension,omitempty"`
}

// Series struct is containing all the series of displayed chart
type Series struct {
	Name       string `json:"name"`
	Categories string `json:"categories,omitempty"`
	Values     string `json:"values"`
}

// Format struct is containing all the chart format
type Format struct {
	XScale          float32 `json:"x_scale,omitempty"`
	YScale          float32 `json:"y_scale,omitempty"`
	XOffset         float32 `json:"x_offset,omitempty"`
	YOffset         float32 `json:"y_offset,omitempty"`
	PrintObj        bool    `json:"print_obj,omitempty"`
	LockAspectRatio bool    `json:"lock_aspect_ratio,omitempty"`
	Locked          bool    `json:"locked,omitempty"`
}

// Legend struct is containing chart legend
type Legend struct {
	Position      string `json:"position,omitempty"`
	ShowLegendKey bool   `json:"show_legend_key,omitempty"`
}

// Title struct is containing chart title
type Title struct {
	Name string `json:"name,omitempty"`
}

// Dimension struct is the dimension of the chart
type Dimension struct {
	Height int `json:"height,omitempty"`
	Width  int `json:"width,omitempty"`
}

// PlotArea struct is containing all options for chart
type PlotArea struct {
	ShowBubbleSize  bool `json:"show_bubble_size,omitempty"`
	ShowCatName     bool `json:"show_cat_name,omitempty"`
	ShowLeaderLines bool `json:"show_leader_lines,omitempty"`
	ShowPercent     bool `json:"show_percent,omitempty"`
	ShowSeriesName  bool `json:"show_series_name,omitempty"`
	ShowVal         bool `json:"show_val,omitempty"`
}
