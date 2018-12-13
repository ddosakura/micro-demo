import { Table, Popconfirm, Button } from 'antd';

const ProductList = ({ onDelete, products }) => {
  const columns = [{
    title: 'Name',
    dataIndex: 'name',
  }, {
    title: 'Actions',
    render: (text, record, index) => {
      // console.log(text, record, index)
      return (
        <Popconfirm title="Delete?" onConfirm={() => onDelete(record.id)}>
    <Button>Delete</Button>
      </Popconfirm>
    );
    },
  }];
  return (
    <Table
      dataSource={products}
      columns={columns}
      rowKey={record => record.id}
    />
  );
};

export default ProductList;

