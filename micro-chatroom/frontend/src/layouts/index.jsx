import styles from './index.css';
import { connect } from 'dva';

function BasicLayout(props) {
  const { client } = props
  const { pc } = client
  return (
    <div className={styles.normal}>
      {pc?<h1 className={styles.title}>Yay! Welcome to umi!</h1>:/*<h1 className={styles.title}>use in moblie</h1>*/<div />}
      {props.children}
    </div>
  );
}
export default connect(({
  client,
}) => ({
  client,
}))(BasicLayout);
