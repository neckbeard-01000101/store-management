import React from "react";
import axios from "axios";
async function handleSubmit(e) {
    e.preventDefault();
    const formData = new FormData(e.target);
    const data = {
        "order-num": formData.get("order-number"),
        "order-state": formData.get("order-state"),
        "customer-name": formData.get("customer-name"),
        "customer-city": formData.get("customer-city"),
        "customer-phone": formData.get("customer-phone"),
        "seller-name": formData.get("seller-name"),
        "total-cost": formData.get("total-cost"),
        "seller-profit": formData.get("seller-profit"),
    };
    try {
        const response = await axios.post(
            "http://localhost:8000/send",
            JSON.stringify(data),
            {
                headers: {
                    "Content-Type": "orders/json",
                },
            },
        );
    } catch (err) {
        console.error(err);
    }
}
function handleNumericInput(
    e,
    maxLength = Infinity,
    min = -Infinity,
    max = Infinity,
) {
    let value = e.target.value;
    if (isNaN(parseInt(value.slice(-1))))
        e.target.value = e.target.value.slice(0, -1);
    if (!isNaN(max) && !isNaN(min)) {
        if (
            parseInt(e.target.value) > max ||
            parseInt(e.target.value) * 1 < min
        ) {
            e.target.value = e.target.value.slice(0, -1);
        }

        if (!isNaN(maxLength)) {
            if (e.target.value.length > maxLength) {
                e.target.value = e.target.value.slice(0, -1);
            }
        }
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
                labelName: "Order state",
                placeHolder: "Order state",
                inputType: "text",
                id: "order-state",
                isNumeric: false,
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
                labelName: "Total cost",
                placeHolder: "Total cost",
                inputType: "text",
                id: "total-cost",
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
                labelName: "Color",
                id: "color",
                options: ["Black", "White"],
            },
            {
                labelName: "Clothes type",
                id: "clothes-type",
                options: ["Hoodie", "Slim T-shirt", "Half T-shirt"],
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
            <button>Submit</button>
        </form>
    );
}

function FormElement(props) {
    const {
        labelName,
        id,
        options = null, // array of options
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
                onChange={
                    isNumeric
                        ? (e) => {
                              handleNumericInput(e);
                          }
                        : null
                }
                type={inputType}
                id={id}
                placeholder={placeHolder}
                name={id}
            />
        </div>
    );
}

export default MainForm;
