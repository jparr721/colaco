# Colaco

## User Story

**Title:** Virtual Vending Machine Development for ColaCo

**As a** software developer at ColaCo,

**I want to** build a user-friendly and visually appealing virtual vending machine webpage,

**So that** customers can easily purchase and download virtual sodas while having an experience similar to using a traditional vending machine.

### Acceptance Criteria

1. **Interactive UI:**
   - The webpage must mimic the look and feel of a traditional soda vending machine.
   - Customers should be able to select and purchase virtual sodas.

2. **Soda Varieties and Availability:**
   - Initially, the vending machine will offer 4 varieties of virtual sodas (Fizz, Pop, Cola, Mega Pop).
   - Each variety has its own cost and limited availability.

3. **Dynamic Pricing and Promotions:**
   - Prices must be adjustable by the sales team based on product performance.
   - The system should allow for price changes during promotional periods as requested by the marketing team.

4. **Purchase and Download:**
   - Upon purchase, a JSON file representing the selected soda should be downloadable.
   - The system must generate the correct soda file based on the customer's selection.

5. **Representation of Money:**
   - The machine should have a virtual representation of money for transactions.

6. **Admin API:**
   - An HTTP API for admins to check the vending machine's status and restock it.
   - The API should consider the limited availability of virtual sodas for restocking.

7. **Price Update Feature:**
   - An easy way for admins to update soda prices.
  
## Deliverable
Please go do `doc/backend.md` to see how to get this application running as well as a full layout of the system and design considerations.
