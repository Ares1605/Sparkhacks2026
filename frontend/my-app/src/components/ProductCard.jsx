export default function ProductCard({ name, desc, price }) {
    return (
        <div className="productCard">
            <h3>{name}</h3>
            <p>{desc}</p>
            <p>{price}</p>
        </div>
    );
}