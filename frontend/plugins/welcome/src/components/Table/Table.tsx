import React, { useState, useEffect, FC } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import AutorenewIcon from '@material-ui/icons/Autorenew';
import { Content, Page, pageTheme,Header } from '@backstage/core';
import { Typography,Grid, } from '@material-ui/core';

import { DefaultApi } from '../../api/apis';
import { EntPayment } from '../../api/models/EntPayment';

const useStyles = makeStyles({
    table: {
      minWidth: 650,
    },
   });
    const ShowTable: FC<{}> = () => {
    const classes = useStyles();
    const api = new DefaultApi();
    const [loading, setLoading] = useState(true);
    const [payments, setPayments ] = useState<EntPayment[]>([]);
    useEffect(() =>{
          const getPayment = async () => {
            const res = await api.listPayment({ limit: 10, offset: 0 });
            setLoading(false);
            setPayments(res);
            console.log(res);
          };
          getPayment();
    },[loading]);
    const refreshPage = ()=>{
      window.location.reload();
    }
    
    return (
      <Page theme={pageTheme.home}>
        <Header
        title={`History Payment`}
      >
      </Header>
      <Content>
        <Grid container >
        <Grid item xs = {6} sm = {6} justify="space-evenly" >
            <Typography variant="h3" gutterBottom >
              ประวัติการชำระ   
              <AutorenewIcon onClick={refreshPage}></AutorenewIcon>
              </Typography>
              <br/>
 
          </Grid>
        <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
         <TableCell align="center">เลขที่บิล</TableCell>
         <TableCell align="center">จำนวนเงิน</TableCell>
         <TableCell align="center">รูปแบบการชำระ</TableCell>
         <TableCell align="center">สถานะการชำระ</TableCell>
         <TableCell align="center">วันเวลาที่แจ้งชำระ</TableCell>
         </TableRow>
         </TableHead>
         <TableBody>
                 {payments.map(item =>(
                <TableRow key={item.id}>
                    <TableCell align="center">{item.edges?.bill?.id}</TableCell>
                    <TableCell align="center">{item.edges?.bill?.amount}</TableCell>
                    <TableCell align="center">{item.edges?.paytype?.typeInfo}</TableCell>
                    <TableCell align="center">{item.edges?.paymentstatus?.paymentStatus}</TableCell>
                    <TableCell align="center">{item.addDatetime}</TableCell>
                </TableRow>
            ))}
         </TableBody>
         </Table>
        </TableContainer>
   </Grid>
      </Content>
    </Page>
    );
   }
   export default ShowTable;