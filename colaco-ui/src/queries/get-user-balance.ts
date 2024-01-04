interface UserAccountBalanceResponse {
    balance: number;
}

export default async function getUserBalance(id: string): Promise<UserAccountBalanceResponse> {
    const response = await fetch(`http://localhost:8000/v1/users/${id}/balance`);
    const data = await response.json();
    return data;
}