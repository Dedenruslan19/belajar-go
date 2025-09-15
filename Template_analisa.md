## ERD Title: Video Game Store System

1. Entities and Their Attributes:
Entity: Table_Name (e.g., Customers)

## Attributes:

- Field_Name (Attribute) : e.g., CustomerId (PK, AI)
- Field (Datatype)
- etc...

A. Entity: Table_Name

- Attributes:
- Field_Name (Attribute) : e.g., FieldId (PK, AI)
- Field (Datatype)
- etc...

B. Entity: Table_Name

- Attributes:
- Field_Name (Attribute) : e.g., FieldId (PK, AI)
- Field (Datatype)
- etc...

C. Entity: Table_Name

- Attributes:
- Field_Name (Attribute) : e.g., FieldId (PK, AI)
- Field (Datatype)
- etc...

## Relationships:
- Table_Name to Table_Name: (e.g., Customers to Orders)

A. Type: One to Many
- Description: One customer can place many orders, but each order is linked to only one customer.
- Table_Name to Table_Name: (e.g., Orders to Order Details)

B. Type: One to Many
- Description: One order can have multiple order details, but each order detail is linked to only one order.
- Table_Name to Table_Name: (e.g., Games to Order Details)

C. Type: One to Many
- Description: One game can appear in many order details, but each order detail is linked to only one game.

## Integrity Constraints:
- The Price in Games and Orders should be a positive float.
- etc...

## Additional Notes:
- The Order Details table allows the system to handle orders with multiple games.
- etc...