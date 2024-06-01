package excel

import "github.com/xuri/excelize/v2"

var DefaultFont = func() *excelize.Font {
	return &excelize.Font{
		Family: "Century Gothic",
		Size:   10,
		Bold:   false,
	}
}()

var DefaultAlignment = func() *excelize.Alignment {
	return &excelize.Alignment{
		Vertical:   "center",
		Horizontal: "center",
		WrapText:   true,
	}
}()

var DefaultStyle = func() *excelize.Style {
	style := &excelize.Style{}
	style.Font = DefaultFont
	style.Alignment = DefaultAlignment
	return style
}()

var TitleStyle = func() *excelize.Style {
	style := &excelize.Style{}
	style.Font = &excelize.Font{
		Family: "Century Gothic",
		Size:   10,
		Bold:   true,
	}
	style.Alignment = DefaultAlignment
	style.Fill = excelize.Fill{
		Type:    "pattern",
		Pattern: 1,
		Color:   []string{"bdc0bf"},
	}

	return style
}()

var DateStyle = func() *excelize.Style {
	style := &excelize.Style{}
	style.Font = DefaultFont
	style.Alignment = DefaultAlignment
	style.NumFmt = 14

	return style
}()

var FloatStyle = func() *excelize.Style {
	fmt := "[$R$-416] #,##0.00;[RED]-[$R$-416] #,##0.00"
	style := &excelize.Style{}
	style.Font = DefaultFont
	style.Alignment = DefaultAlignment
	style.CustomNumFmt = &fmt

	return style
}()
