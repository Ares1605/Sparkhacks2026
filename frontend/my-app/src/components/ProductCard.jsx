import { useState } from 'react'
import './ProductCard.css'

export default function ProductCard({ name, img, price, isBookmarked, onBookmark }) {
    return (
        <div className="productCard">
            <img src={img} className="product-img"/>
            <h3>{name}</h3>
            <p>{price}</p>

            <button className="add-cart" onClick={onBookmark}>
                {isBookmarked ? "Remove" : "Bookmark"}
            </button>
        </div>
    );
}