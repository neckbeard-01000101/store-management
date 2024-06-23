import React, { useEffect, useState } from "react";
import axios from "axios";

const SERVER_URL = "http://localhost:8000";

export default function Data() {
    const [collections, setCollections] = useState([]);
    const [data, setData] = useState([]);
    const [selectedCollection, setSelectedCollection] = useState("06/2024");
    const calculateAndSetPercentage = (index) => {
        const inputField = document.getElementById(`percentage-input-${index}`);
        if (!inputField) {
            console.error("Input field not found");
            return;
        }

        const totalProfit = data[index]["total-profit"];
        if (!totalProfit) {
            console.error(
                "Invalid item or total-profit not found:",
                data[index],
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
    useEffect(() => {
        const fetchCollections = async () => {
            try {
                const response = await axios.get(`${SERVER_URL}/collections`);
                setCollections(response.data);
                setCollections((prev) =>
                    prev.filter(
                        (collection) => collection !== "months-profits",
                    ),
                );
            } catch (error) {
                console.error("Failed to fetch collections:", error);
            }
        };

        fetchCollections();
    }, []);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get(
                    `${SERVER_URL}/get/${selectedCollection}`,
                );
                setData(response.data);
            } catch (error) {
                console.error("Failed to fetch data:", error);
            }
        };

        if (selectedCollection) {
            fetchData();
        }
    }, [selectedCollection]);

    const handleChange = (event) => {
        setSelectedCollection(event.target.value);
    };

    const handleStateChange = async (index) => {
        if (!window.confirm("Are you sure you want to mark the order as done?"))
            return;
        if (index < 0 || index >= data.length) {
            console.error("Invalid index:", index);
            return;
        }

        const newData = [...data];
        const item = newData[index];
        if (!item) {
            console.error("Item not found at index:", index);
            return;
        }

        if (!("order-state" in item)) {
            console.error("Property 'orderState' not found in item:", item);
            return;
        }
        const collectionName = item["collection-name"];
        const newState = item["order-state"] === "Undone" ? "Done" : "Undone";
        item["order-state"] = newState;
        setData(newData);
        try {
            await axios.post(
                `${SERVER_URL}/toggleOrderState/${item._id}?newState=${newState}&collectionName=${collectionName}`,
            );
            console.log("New state:", newState);
            alert("Order state changed successfully");
        } catch (error) {
            console.error("Failed to update order state:", error);
            alert("There was an error changing the order state");
        }
    };

    return (
        <div className="data-page">
            <label htmlFor="month">Select a month: </label>
            <select
                id="month"
                value={selectedCollection}
                onChange={handleChange}
            >
                {collections.map((collection) => (
                    <option key={collection} value={collection}>
                        {collection}
                    </option>
                ))}
            </select>
            {selectedCollection && (
                <div>
                    <h2>{selectedCollection}</h2>
                    <table>
                        <thead>
                            <tr>
                                <th>Order Number</th>
                                <th>Order State</th>
                                <th>Customer Name</th>
                                <th>Customer City</th>
                                <th>Customer Phone</th>
                                <th>Seller Name</th>
                                <th>Total Cost</th>
                                <th>Seller Profit</th>
                                <th>Delivery Fee</th>
                                <th>Cost of Product</th>
                                <th>Number of pieces</th>
                                <th>Size</th>
                                <th>Color</th>
                                <th>Clothes Type</th>
                                <th>Total profit</th>
                                <th>Profit percentage</th>
                            </tr>
                        </thead>
                        <tbody>
                            {data.map((item, index) => (
                                <tr key={index}>
                                    <td>{item["order-num"]}</td>
                                    <td
                                        style={
                                            item["order-state"] === "Done"
                                                ? {
                                                      backgroundColor:
                                                          "#abffc4",
                                                  }
                                                : null
                                        }
                                    >
                                        {item["order-state"]}

                                        <button
                                            className="state-btn"
                                            onClick={() =>
                                                handleStateChange(index)
                                            }
                                        >
                                            {item["order-state"] === "Undone"
                                                ? "Mark as done"
                                                : "Mark as undone"}
                                        </button>
                                    </td>
                                    <td>{item["customer-name"]}</td>
                                    <td>{item["customer-city"]}</td>
                                    <td>{item["customer-phone"]}</td>
                                    <td>{item["seller-name"]}</td>
                                    <td>{item["total-cost"]}</td>
                                    <td>{item["seller-profit"]}</td>
                                    <td>{item["delivery-fee"]}</td>
                                    <td>{item["cost-of-product"]}</td>
                                    <td>{item["pieces-num"]}</td>
                                    <td>{item.size}</td>
                                    <td>{item.color}</td>
                                    <td>{item["clothes-type"]}</td>
                                    <td>{item["total-profit"]}</td>
                                    <td>
                                        <input
                                            type="text"
                                            id={`percentage-input-${index}`}
                                            placeholder="Enter percentage"
                                        />
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
            )}
        </div>
    );
}
