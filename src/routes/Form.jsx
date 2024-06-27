import React from "react";
import axios from "axios";

async function handleSubmit(e) {
    e.preventDefault();
    if (!confirm("Are you sure you want to submit the data?")) return;
    const formData = new FormData(e.target);
    const data = {
        "order-num": parseInt(formData.get("order-number"), 10),
        "order-state": formData.get("order-state"),
        "customer-name": formData.get("customer-name"),
        "customer-city": formData.get("customer-city"),
        "customer-phone": parseInt(formData.get("customer-phone"), 10),
        "seller-name": formData.get("seller-name"),
        "total-cost": parseInt(formData.get("total-cost"), 10),
        "seller-profit": parseInt(formData.get("seller-profit"), 10),
        "delivery-fee": parseInt(formData.get("fee"), 10),
        "cost-of-product": parseInt(formData.get("cost-of-product"), 10),
        "pieces-num": parseInt(formData.get("pieces-num"), 10),
        size: formData.get("size"),
        color: formData.get("color"),
        "clothes-type": formData.get("clothes-type"),
    };

    try {
        const response = await axios.post(
            "http://localhost:8000/send",
            JSON.stringify(data),
            {
                headers: {
                    "Content-Type": "application/json",
                },
            },
        );
        console.log(response);
        if (response.status === 200) {
            alert("Data sent successfully");
        }
    } catch (err) {
        console.error(err);
        alert("Error sending data", err);
    }
}

export function HandleNumericInput(
    e,
    maxLength = Infinity,
    min = -Infinity,
    max = Infinity,
) {
    let value = e.target.value;
    let lastChar = value.slice(-1);

    if (isNaN(parseInt(lastChar))) {
        e.target.value = value.slice(0, -1);
        return;
    }
    let numericValue = parseInt(value);

    if (numericValue > max || numericValue < min || value.length > maxLength) {
        e.target.value = value.slice(0, -1);
    }
}

const fields = [
    {
        customer: [
            {
                labelName: "Order number",
                placeHolder: "Order number",
                inputType: "text",
                id: "order-number",
                isNumeric: true,
            },
            {
                labelName: "Customer name",
                placeHolder: "Customer name",
                inputType: "text",
                id: "customer-name",
                isNumeric: false,
            },
            {
                labelName: "Customer phone number",
                placeHolder: "Customer phone number",
                inputType: "text",
                id: "customer-phone",
                isNumeric: true,
            },
            {
                labelName: "Number of pieces",
                placeHolder: "Number of pieces",
                inputType: "text",
                id: "pieces-num",
                isNumeric: true,
            },
            {
                labelName: "Customer city",
                placeHolder: "Customer city",
                inputType: "text",
                id: "customer-city",
                isNumeric: false,
            },
        ],
        seller: [
            {
                labelName: "Seller name",
                placeHolder: "Seller name",
                inputType: "text",
                id: "seller-name",
                isNumeric: false,
            },
            {
                labelName: "Seller profit",
                placeHolder: "Seller profit",
                inputType: "text",
                id: "seller-profit",
                isNumeric: true,
            },
            {
                labelName: "Cost per piece",
                placeHolder: "Cost per piece",
                inputType: "text",
                id: "total-cost",
                isNumeric: true,
            },
            {
                labelName: "Cost of product",
                placeHolder: "Cost of product",
                inputType: "text",
                id: "cost-of-product",
                isNumeric: true,
            },
        ],
        "customer-optins": [
            {
                labelName: "Size",
                id: "size",
                options: ["S", "M", "L", "XL", "XLL", "XLLL", "XLLLL"],
            },
            {
                labelName: "Order state",
                id: "order-state",
                options: ["Undone", "Done"],
            },
            {
                labelName: "Color",
                id: "color",
                options: ["Black", "White"],
            },
            {
                labelName: "Clothes type",
                id: "clothes-type",
                options: ["Hoodie", "Slim T-shirt", "Hav T-shirt"],
            },
            {
                labelName: "Delivery fee",
                id: "fee",
                options: ["5000", "6000"],
            },
        ],
    },
];

function MainForm() {
    return (
        <form className="main-form" onSubmit={(e) => handleSubmit(e)}>
            <div className="fields customer-fields">
                {fields[0].customer.map((field) => (
                    <FormElement
                        labelName={field.labelName}
                        id={field.id}
                        key={field.id}
                        placeHolder={field.placeHolder}
                        inputType={field.inputType}
                        isNumeric={field.isNumeric}
                    />
                ))}
                {fields[0]["customer-optins"].map((option) => (
                    <FormElement
                        labelName={option.labelName}
                        id={option.id}
                        key={option.id}
                        options={option.options}
                        isSelect={true}
                    />
                ))}
            </div>
            <div className="fields seller-fields">
                {fields[0].seller.map((field) => (
                    <FormElement
                        labelName={field.labelName}
                        id={field.id}
                        key={field.id}
                        placeHolder={field.placeHolder}
                        inputType={field.inputType}
                        isNumeric={field.isNumeric}
                    />
                ))}
            </div>
            <button type="submit">Submit</button>
        </form>
    );
}

function FormElement(props) {
    const {
        labelName,
        id,
        options = null,
        isSelect = false,
        placeHolder = null,
        inputType = null,
        isNumeric = false,
    } = props;

    if (isSelect && options) {
        return (
            <div className="form-element">
                <label htmlFor={id}>{labelName}</label>
                <select name={id} id={id}>
                    {options.map((option) => (
                        <option key={option} value={option}>
                            {option}
                        </option>
                    ))}
                </select>
            </div>
        );
    }

    return (
        <div className="form-element">
            <label htmlFor={id}>{labelName}</label>
            <input
                onChange={isNumeric ? (e) => HandleNumericInput(e) : null}
                type={inputType}
                id={id}
                placeholder={placeHolder}
                name={id}
            />
        </div>
    );
}
export default MainForm;
