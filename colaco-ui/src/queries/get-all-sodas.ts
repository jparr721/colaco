export interface Soda {
	id: string;
	product_name: string;
	description: string;
	cost: number;
	current_quantity: number;
	max_quantity: number;
	created_at: string;
	updated_at: string;
}

export default async function getAllSodas(): Promise<Soda[]> {
    const response = await fetch(`http://localhost:8000/v1/sodas`);
    const data = await response.json();
    return data;
}