import { useState, useEffect } from 'react'
import '../App.css'
import ProductCard from "../components/ProductCard"
import Sidebar from '../components/SideBar';

function Home() {
  
  const products = [
    {
      id: 1,
      name: "IPhone 16 Pro",
      img: "https://via.placeholder.com/300x200?text=iPhone",
      price: "$899.99"
    },
    {
      id: 2,
      name: "Red Bull",
      img: "https://via.placeholder.com/300x200?text=iPhone",
      price: "$2.99"
    },
    {
      id: 3,
      name: "Lamp",
      img: "https://via.placeholder.com/300x200?text=iPhone",
      price: "$44.99"
    },
    {
      id: 4,
      name: "Bag of Chips",
      img: "https://via.placeholder.com/300x200?text=iPhone",
      price: "$4.99"
    },
    {
      id: 5,
      name: "A Chair",
      img: "https://via.placeholder.com/300x200?text=iPhone",
      price: "$149.99"
    }
  ]

  const getInitialBookmarks = () => {
    const stored = localStorage.getItem("bookmark");
    return stored ? JSON.parse(stored) : [];
  }

  const [bookmark, setBookmarks] = useState(getInitialBookmarks);

  useEffect(() => {
    localStorage.setItem("bookmark", JSON.stringify(bookmark));
  }, [bookmark]);

  const addToBookmark = (product) => {
    setBookmarks((prev) => {
      if (prev.find((item) => item.id === product.id)) {
        return prev;
      }
      return [...prev, product];
    });
  };

  const removeFromBookmark = (id) => {
    setBookmarks((prev) => prev.filter((item) => item.id !== id));
  }

  return (
    <div className="layout">
      <div className="main">
        <div className="mainBox">
          <div className="product-slider">
            {products.map((product) => (
              <ProductCard
                key={product.id}
                name={product.name}
                img={product.img}
                price={product.price}
                isBookmarked={false}
                onBookmark={() => addToBookmark(product)}
              />
            ))}
          </div>
        </div>

        <div className="subBox">
          <h2>Your Bookmarks</h2>
          <div className="bookmark-container">
            {bookmark.length === 0 ? (
              <p>No items added yet.</p>
            ) : (
              <div className="product-slider">
                {bookmark.map((item) => (
                  <ProductCard
                    key={item.id}
                    name={item.name}
                    img={item.img}
                    price={item.price}
                    isBookmarked={true}
                    onBookmark={() => removeFromBookmark(item.id)}
                  />
                ))}
              </div>
            )}
          </div>
        </div>
      </div>

      <Sidebar/>
    </div>
  );
}

export default Home;