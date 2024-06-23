import React, { useEffect, useState } from "react";
import axios from "axios";

function Profits() {
    const [profits, setProfits] = useState([]);

    useEffect(() => {
        const fetchMonths = async () => {
            try {
                const response = await axios.get(
                    "http://localhost:8000/get/months-profits",
                );
                setProfits(response.data);
            } catch (error) {
                console.error("Faild to fetch profits:", error);
            }
        };

        fetchMonths();
    }, []);

    return (
        <div className="profits">
            <h2>Monthly Profits</h2>
            <table>
                <thead>
                    <tr>
                        <th>Month</th>
                        <th>Profit</th>
                    </tr>
                </thead>
                <tbody>
                    {profits.map((profit, index) => (
                        <tr key={index}>
                            <td>{profit.month}</td>
                            <td>{profit.profit}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
}

export default Profits;

