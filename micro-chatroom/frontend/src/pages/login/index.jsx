import styles from './index.css';
import { connect } from 'dva';
import { Form, Icon, Input, Button, Checkbox } from 'antd';
import { Switch } from 'antd-mobile';

const FormItem = Form.Item;

class NormalLoginForm extends React.Component {
  handleSubmit = (e) => {
    e.preventDefault();
    this.props.form.validateFields((err, values) => {
      if (!err) {
        this.props.handleLogin(values)
      }
    });
  }

  render() {
    const { getFieldDecorator } = this.props.form;
    return (
      <Form onSubmit={this.handleSubmit} className={styles.form}>
        <FormItem>
          {getFieldDecorator('userName', {
            rules: [{ required: true, message: 'Please input your username!' }],
          })(
            <Input prefix={<Icon type="user" style={{ color: 'rgba(0,0,0,.25)' }} />} placeholder="Username" />
          )}
        </FormItem>
        <FormItem>
          {getFieldDecorator('password', {
            rules: [{ required: true, message: 'Please input your Password!' }],
          })(
            <Input prefix={<Icon type="lock" style={{ color: 'rgba(0,0,0,.25)' }} />} type="password" placeholder="Password" />
          )}
        </FormItem>
        <FormItem>
          {getFieldDecorator('remember', {
            valuePropName: 'checked',
            initialValue: true,
          })(
            <Checkbox>Remember me</Checkbox>
          )}
          <a className={styles.forgot} href="">Forgot password</a>
          <Button type="primary" htmlType="submit" className={styles.button}>
            Log in
          </Button>
          Or <a href="">register now!</a>
        </FormItem>
      </Form>
    );
  }
}

const WrappedNormalLoginForm = Form.create()(NormalLoginForm);

export default connect(({
  client,
}) => ({
  client,
}))(function ({ client, dispatch }) {
  function handleClick() {
    dispatch({
      type: 'client/change',
      pc: !pc,
    })
  }
  const { pc } = client
  if (pc) {
    return (<div>
      <Switch
        checked={pc}
        onChange={handleClick}
      />
      <WrappedNormalLoginForm handleLogin={(values) => {

        console.log('Received values of form: ', values);
      }} /></div>
    )
  }
  return (<div>
    <Switch
      checked={pc}
      onChange={handleClick}
    />
  </div>)
});
