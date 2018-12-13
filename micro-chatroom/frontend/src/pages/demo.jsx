import styles from './demo.css';
import axios from 'axios';
import {
  connect
} from 'dva';
import {
  Button
} from 'antd';
import ProductList from '../components/ProductList';

const Products = ({
  dispatch,
  products
}) => {
  function handleDelete(id) {
    dispatch({
      type: 'products/delete',
      payload: id,
    });
  }
  function handleFresh(e) {
    // console.log(e)
    axios.get('/api/products')
      .then((res) => {
        dispatch({
          type: 'products/fresh',
          payload: res.data,
        });
      })
  }
  return (<div>
    <Button type="primary" onClick={handleFresh}>fresh</Button>
    <h2>List of Products</h2>
    <ProductList onDelete={handleDelete} products={products} />
    <iframe style={{border:0,width:"100%",height:630}} src="/search?wd=123"/>
  </div>
  );
};

export default connect(({
  products
}) => ({
  products,
}))(Products);
