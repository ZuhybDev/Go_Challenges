Write a function processMarketInventory that takes inventory, sales and returns a reversed slice of sold-out items.

The function processes vendor inventory and sales orders to track which items are completely sold out, then returns those items in reverse order of their original appearance.

Logic:

Parse inventory entries in format "vendor:item:quantity" to build stock tracking
Process each sales order, updating quantities (use continue to skip invalid orders)
Collect items that reach zero quantity after all sales
Return sold-out items in reverse order of their original inventory position
Conditions:

Skip sales orders that don't match existing inventory items
Skip sales orders with invalid quantity format
Only include items that become completely sold out (quantity = 0)
Parameters:

inventory ([]string): Vendor inventory in format "vendor:item:quantity"
sales ([]string): Sales orders in format "item:soldQuantity"
Returns: Slice of sold-out item names in reverse order. Format: ["item3", "item1"]
