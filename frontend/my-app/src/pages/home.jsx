import { useState, useEffect } from 'react'
import '../App.css'
import ProductCard from "../components/ProductCard"
import Sidebar from '../components/SideBar';
import ChatBox from '../components/ChatBox';


function Home({ api }) {
  const [products, setProducts] = useState([])

  useEffect(() => {
    api.getRecommendations().
      then(({ data }) => {
        data = JSON.parse(data).data
        setProducts(data.map((d, i) => ({
          id: i,
          name: d.name,
          img: api.resolveImagePath2(d.image_path),
          price: 0,
        })))
      })
  }, [api])

  const getInitialBookmarks = () => {
    const stored = localStorage.getItem("bookmark");
    return stored ? JSON.parse(stored) : [];
  }

  const [bookmark, setBookmarks] = useState(getInitialBookmarks);
  const [notify, setNotify] = useState('');

  useEffect(() => {
    localStorage.setItem("bookmark", JSON.stringify(bookmark));
  }, [bookmark]);

  const showNotification = (message) => {
    setNotify(message);
    setTimeout(() => setNotify(''), 2000);
  };

  const addToBookmark = (product) => {
    setBookmarks((prev) => {
      if (prev.find((item) => item.id === product.id)) {
        return prev;
      }
      showNotification(`${product.name} added to Bookmarks`);
      return [...prev, product];
    });
  };

  const removeFromBookmark = (id) => {
    const product = bookmark.find((item) => item.id === id);
    if (product) {
      showNotification(`${product.name} removed from Bookmarks`);
    }
    setBookmarks((prev) => prev.filter((item) => item.id !== id));
  }

  return (
    <div className="layout">
      <div className="main">
        <div className="mainBox">
          <h2>Recommendations</h2>
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

        <ChatBox />

      </div>

      <Sidebar/>

      {notify && (
        <div className="notification">
          {notify}
        </div>
      )}
    </div>
  );
}

export default Home;
