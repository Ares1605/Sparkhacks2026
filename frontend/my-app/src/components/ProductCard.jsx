import './ProductCard.css'

export default function ProductCard({ name, img, price, isBookmarked, onBookmark }) {
    return (
        <div className="productCard">
            <img src={img} className="product-img" alt={name} />
            
            <h3>{name}</h3>
            
            <p>{price}</p>

            <button 
                className={`add-cart ${isBookmarked ? 'is-active' : ''}`} 
                onClick={onBookmark}
                style={{
                    backgroundColor: isBookmarked ? '#f3f4f6' : '',
                    color: isBookmarked ? '#1f2937' : ''
                }}
            >
                {isBookmarked ? "Remove" : "Bookmark"}
            </button>
        </div>
    );
}
