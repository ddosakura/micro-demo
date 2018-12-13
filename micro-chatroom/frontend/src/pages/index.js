import styles from './index.css';
import logo from '../assets/yay.jpg';
import Link from 'umi/link';

export default function() {
  return (
    <div className={styles.normal}>
      <div className={styles.welcome} />
      <ul className={styles.list}>
        <li><Link to="/demo">Getting Started</Link></li>
        <li><Link to="/login">Login</Link></li>
      </ul>
      <img src={logo} alt="Logo"/>
      <ul className={styles.list}>
        <li>To get started, edit <code>src/pages/index.js</code> and save to reload.</li>
      </ul>
    </div>
  );
}
