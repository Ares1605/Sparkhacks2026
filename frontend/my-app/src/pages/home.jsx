import '../App.css'
import ProductCard from "../components/ProductCard"
import { Link } from "react-router-dom";

function Home() {
  const products = [
    {
      id: 1,
      name: "IPhone 16 Pro",
      desc: "A Cell Phone",
      price: "$899.99"
    },
    {
      id: 2,
      name: "A Chair",
      desc: "A Place to sit on",
      price: "$149.99"
    }
  ]

  return (
    <div className="layout">
      <div className="main">
        <div className="mainBox">
          <h1>Main Placeholder</h1>
          <p>yap yap yap yap yap</p>

          <div className="productGrid">
            {products.map((product) => (
              <ProductCard
                key={product.id}
                name={product.name}
                description={product.desc}
                price={product.price}
              />
            ))}
          </div>
        </div>

        <div className="subBox">
          <h2>Bottom Box</h2>
          <p>Additional content</p>
        </div>
      </div>

      <div className="sidebar">
        <h2>Sidebar</h2>
        <p>Navigation</p>
        <Link to="/sync" style={{ color: '#60a5fa', textDecoration: 'none' }}> Sync </Link>
      </div>
    </div>
  );
}

export default Home;