import React, { useEffect, useState } from "react";
import axios from "axios";
import { HandleNumericInput } from "./Form";
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

    const calculateAndSetPercentage = (index) => {
        const inputField = document.getElementById(`percentage-month-${index}`);
        if (!inputField) {
            console.error("Input field not found");
            return;
        }

        const totalProfit = profits[index]["profit"];
        if (!totalProfit) {
            console.error(
                "Invalid item or total-profit not found:",
                profits[index],
            );
            return;
        }

        const userPercentage = parseFloat(inputField.value);
        if (isNaN(userPercentage)) {
            console.error("Invalid percentage input:", inputField.value);
            return;
        }

        const calculatedPercentage = (userPercentage / 100) * totalProfit;
        inputField.value = calculatedPercentage.toString();
    };
    return (
        <div className="profits">
            <h2>Monthly Profits</h2>
            <table>
                <thead>
                    <tr>
                        <th>Month</th>
                        <th>Profit</th>
                        <th>Profit percentage</th>
                    </tr>
                </thead>
                <tbody>
                    {profits?.map((profit, index) => (
                        <tr key={index}>
                            <td>{profit.month}</td>
                            <td>{profit.profit}</td>
                            <td>
                                <input
                                    type="text"
                                    id={`percentage-month-${index}`}
                                    placeholder="Enter percentage"
                                    onChange={(e) =>
                                        HandleNumericInput(e, 3, 0, 100)
                                    }
                                />
                                <br />
                                <button
                                    onClick={() =>
                                        calculateAndSetPercentage(index)
                                    }
                                >
                                    Calculate
                                </button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
}

export default Profits;
