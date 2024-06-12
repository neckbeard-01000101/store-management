import React from "react";

const fields = [
    {
        customer: [
            {
                labelName: "Order number",
                placeHolder: "Order number",
                inputType: "text",
                id: "Order-number",
            },
            {
                labelName: "Order state",
                placeHolder: "Order state",
                inputType: "text",
                id: "order-state",
            },
            {
                labelName: "Customer name",
                placeHolder: "Customer name",
                inputType: "text",
                id: "customer-name",
            },
            {
                labelName: "Customer city",
                placeHolder: "Customer city",
                inputType: "text",
                id: "customer-city",
            },
        ],
        seller: [
            {
                labelName: "Seller name",
                placeHolder: "Seller name",
                inputType: "text",
                id: "seller-name",
            },
            {
                labelName: "Seller profit",
                placeHolder: "Seller profit",
                inputType: "text",
                id: "seller-profit",
            },
            {
                labelName: "Total cost",
                placeHolder: "Total cost",
                inputType: "text",
                id: "total-cost",
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
        <form className="main-form">
            <div className="fields customer-fields">
                {fields[0].customer.map((field) => (
                    <FormElement
                        labelName={field.labelName}
                        id={field.id}
                        key={field.id}
                        placeHolder={field.placeHolder}
                        inputType={field.inputType}
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
    } = props;

    if (isSelect && options) {
        return (
            <div className="form-element">
                <label htmlFor={id}>{labelName}</label>
                <select name={labelName} id={id}>
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
            <input type={inputType} id={id} placeholder={placeHolder} />
        </div>
    );
}

export default MainForm;
