import React, { useEffect } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import {Grid,InputLabel,MenuItem,Button,TextField,FormControl,Select, Link, Typography} from '@material-ui/core';
import { Link as RouterLink } from 'react-router-dom';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from'../../api/apis';
import { EntBill } from '../../api/models/EntBill'; // import interface Bill
import { EntPaymentStatus } from '../../api/models/EntPaymentStatus'; // import interface PaymentStatus
import { EntPayType } from '../../api/models/EntPayType'; // import interface PayType

const useStyles = makeStyles((theme: Theme) =>
    createStyles({
      button: {
        display: 'block',
        marginTop: theme.spacing(2),
      },
      formControl: {
          width: 800,
        },
        selectEmpty: {
          marginTop: theme.spacing(2),
        },
    }),
  );

export  default  function Create() {
    const classes = useStyles();
    const http = new DefaultApi();

    const [status, setStatus] = React.useState(false);
    const [alert, setAlert] = React.useState(true);

    const [paymentstatuss, setPaymentStatuss] = React.useState<EntPaymentStatus[]>([]);
    const [paytypes, setPayTypes] = React.useState<EntPayType[]>([]);
    const [bills, setBills] = React.useState<EntBill[]>([]);
    const [datetime, setDatetime] = React.useState(String);
    const [email, setEmail] = React.useState(String);
    const [billid, setbillId] = React.useState(Number);
    const [paytypeid, setpaytypeId] = React.useState(Number);
    const [paymentstatusid, setpaymentstatusId] = React.useState(Number);
    const [loading, setLoading] = React.useState(true);

    useEffect(() => {
        const getBill = async () => {
            const res = await http.listBill({ limit: 100, offset: 0 });
            setLoading(false);
            setBills(res);
          };
          const getPaymentStatus = async () => {
            const res = await http.listPaymentstatus({ limit: 2, offset: 0 });
            setLoading(false);
            setPaymentStatuss(res);
          };
        const getPayType = async () => {
            const res = await http.listPaytype({ limit: 2, offset: 0 });
            setLoading(false);
            setPayTypes(res);
        };
        getBill();
        getPaymentStatus();
        getPayType();
    }, [loading]);

    
  const BillhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setbillId(event.target.value as number);
  };
  const PaytypehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setpaytypeId(event.target.value as number);
  };
  const PaymentstatushandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setpaymentstatusId(event.target.value as number);
  };
  const handleDatetimeChange = (event: any) => {
    setDatetime(event.target.value as string);
  };
  const UserhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setEmail(event.target.value as string);
  };
  const CreatePayment = async () => {
    const payment = {
      bill : billid,
      payType: paytypeid ,
      paymentStatus : paymentstatusid,
      addDatetime : datetime + ":00+07:00"

    };
    const bill = {
      id : billid,
      billStatus : 2
    }
    console.log(payment);
    await http.updateBill({id:billid,bill:bill});
    const res: any = await http.createPayment({ payment : payment });
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
    } else {
      setAlert(false);
    }
    setTimeout(() => {
      setStatus(false);
    }, 1000);
  };
  return(
    <div>
          {status ? (
             <div>
             {alert ? (
               <Alert severity="success">
                 บันทึกการชำระสำเร็จ
               </Alert>
             ) : (
                 <Alert severity="warning" style={{ marginTop: 20 }}>
                   มีข้อผิดพลาด โปรดลองอีกครั้ง
                 </Alert>
               )}
           </div>
          ):null}
      <Grid container spacing={3}>
      <Grid item xs={12}>
        <Typography align = "center"> 
          <TextField 
          className={classes.formControl}
          required id="user"
          label="email" 
          value={email}
          onChange={UserhandleChange}/>
        </Typography>
      </Grid>
      <Grid item xs={12}>
      <Typography align = "center">
        <FormControl className={classes.formControl} >
          <InputLabel>Bill ID</InputLabel>
         <Select
            name="bill"
            value={billid}
            onChange={BillhandleChange}
         >
            {bills.filter(b => b.edges?.resident?.email===email).filter(bs => bs.edges?.billstatus?.billStatus === "Unpaid").map(item => (
              <MenuItem value={item.id}>{item.id}</MenuItem>                 
            ))}
         </Select>
        </FormControl>
      </Typography>   
      </Grid>
      <Grid item xs={12}>
      <Typography align = "center">
        <FormControl className={classes.formControl}>
          <InputLabel>PayType ID</InputLabel>
          <Select
          name="paytype"
          value={paytypeid}
          onChange={PaytypehandleChange}
          >
          {paytypes.map(item => {
                return (
                  <MenuItem value={item.id}>
                    {item.typeInfo}
                  </MenuItem>
                );
              })}
          </Select>
        </FormControl> 
      </Typography>
      </Grid>
      <Grid item xs={12}>
      <Typography align = "center">
        <FormControl className={classes.formControl}>
          <InputLabel>PaymentStatus ID</InputLabel>
          <Select
              name="paymentstatus"
              value={paymentstatusid}
              onChange={PaymentstatushandleChange}
          >
              {paymentstatuss.map(item => {
                return (
                  <MenuItem value={item.id}>
                    {item.paymentStatus}
                  </MenuItem>
                );
              })}
          </Select>
        </FormControl> 
        </Typography>
      </Grid>
      <Grid item xs={12}>
      <Typography align = "center">
            <TextField
              className={classes.formControl}
              id="datetime"
              label="ว/ด/ป "
              type="datetime-local"
              value={datetime}
              onChange={handleDatetimeChange}
              InputLabelProps={{
                shrink: true,
              }}
            />
            </Typography>
      </Grid>
      <br/>
      <Grid item xs={6} >
      <Typography align = "center">
        <Button
            onClick={() => {
              CreatePayment();
            }}
            variant="contained"
            color = "primary"
          >
            บันทึกการชำระ
        </Button>
        </Typography>
      </Grid>
      <Grid item xs={6}>
      <Typography align = "center">
        <Link component = {RouterLink} to ="/Home">
          <Button variant="contained">
            กลับ
          </Button>
        </Link>
        </Typography>
      </Grid>
      </Grid>
    </div>
);
}

