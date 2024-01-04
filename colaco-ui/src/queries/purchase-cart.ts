import { DropdownItem } from "../components/dropdown-button";

interface ReceiptLine {
    id: string;
    product_name: string;
    price: number;
}

export interface Receipt {
    items: ReceiptLine[];
    total: number;
}

export default async function purchaseCart(id: string, items: DropdownItem[]): Promise<Receipt> {
    console.log(`Purchasing cart ${id} with items:`, items);
    // Submit a PUT to the user withdraw endpoint for each individual item in the cart.
    const processedItems: DropdownItem[] = [];
    for (const item of items) {
        try {
            await fetch(`http://localhost:8000/v1/users/me/withdraw`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    'x-auth-token': id,
                },
                body: JSON.stringify({
                    amount: item.price,
                })
            });

            await fetch(`http://localhost:8000/v1/sodas/${item.id}/sell`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                    'x-auth-token': id,
                },
                body: JSON.stringify({amount: 1})
            });
            processedItems.push(item);
        } catch {
            // Revert the purchase if any of the items fail.
            for (const revertItem of processedItems) {
                await fetch(`http://localhost:8000/v1/users/me/deposit`, {
                    method: 'PUT',
                    headers: {
                        'Content-Type': 'application/json',
                        'x-auth-token': id,
                    },
                    body: JSON.stringify({
                        amount: revertItem.price,
                    })
                });

                // If the user isn't an admin this melts, lol.
                await fetch(`http://localhost:8000/v1/sodas/${revertItem.id}/restock`, {
                    method: 'PUT',
                    headers: {
                        'Content-Type': 'application/json',
                        'x-auth-token': id,
                    },
                    body: JSON.stringify({amount: 1})
                });
            }

            throw new Error(`Failed to purchase item ${item.id} ${item.product_name}`);
        }
    }

    return {
        items: processedItems,
        total: processedItems.reduce((acc, item) => acc + item.price, 0),
    };
}