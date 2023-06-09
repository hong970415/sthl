package constants

type contextKey string

// Product Status Type
type productStatusType struct {
	Initiated  string
	Active     string
	OutOfStock string
	Inactive   string
}

func (p productStatusType) GetList() []string {
	return []string{
		p.Initiated,
		p.Active,
		p.OutOfStock,
		p.Inactive,
	}
}

// Order Status Type
type orderStatusType struct {
	Initiated string
	Confirmed string
	Shipping  string
	Canceled  string
	Completed string
}

func (o orderStatusType) GetList() []string {
	return []string{
		o.Initiated,
		o.Confirmed,
		o.Shipping,
		o.Canceled,
		o.Completed,
	}
}

// Payment Status Type
type paymentStatusType struct {
	Pending         string
	Fail            string
	Paid            string
	Refunded        string
	PartialRefunded string
	NoRefund        string
	Voided          string
}

func (p paymentStatusType) GetList() []string {
	return []string{
		p.Pending,
		p.Fail,
		p.Paid,
		p.Refunded,
		p.PartialRefunded,
		p.NoRefund,
		p.Voided,
	}
}

// Payment MethodType
type paymentMethodType struct {
	// Cash         string
	Card string
}

func (p paymentMethodType) GetList() []string {
	return []string{
		p.Card,
	}
}

// Delivery Status Type
type deliveryStatusType struct {
	Pending   string
	Shipping  string
	Completed string
}

func (d deliveryStatusType) GetList() []string {
	return []string{
		d.Pending,
		d.Shipping,
		d.Completed,
	}
}
