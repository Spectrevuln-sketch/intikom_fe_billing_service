"use client";

import { useEffect, useState } from "react";

export default function Home() {
  const [data, setData] = useState<any>(null);
  const [page, setPage] = useState(1);

  const limit = 5;

  const fetchData = async () => {
    try {
      const res = await fetch(
        `http://localhost:8080/accounts/user_1/transactions?page=${page}&limit=${limit}`
      );
      const json = await res.json();
      setData(json);
    } catch (err) {
      console.error("fetch error:", err);
    }
  };

  useEffect(() => {
    fetchData();

    const interval = setInterval(fetchData, 3000);
    return () => clearInterval(interval);
  }, [page]);

  if (!data) return <p>Loading...</p>;

  return (
    <div style={{ padding: 20, fontFamily: "Arial" }}>
      <h1>Transaction Monitoring</h1>

      <h2>Account: user_1</h2>
      <h3>Balance: {data.balance}</h3>

      <p>
        Page: {data.page} | Total: {data.total}
      </p>

      <h4>Transactions</h4>
      <ul>
        {data.data.map((trx: any, i: number) => (
          <li key={i}>
            {trx.type} - {trx.amount} ({trx.reference_id})
          </li>
        ))}
      </ul>

      <div style={{ marginTop: 20 }}>
        <button
          onClick={() => setPage((p) => Math.max(p - 1, 1))}
          disabled={page === 1}
        >
          Prev
        </button>

        <span style={{ margin: "0 10px" }}>Page {page}</span>

        <button
          onClick={() =>
            setPage((p) =>
              p * limit < data.total ? p + 1 : p
            )
          }
        >
          Next
        </button>
      </div>
    </div>
  );
}