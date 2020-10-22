import React, { FC } from 'react';
import { Typography,Grid, } from '@material-ui/core';
import { Content, Header, Page, pageTheme, } from '@backstage/core';
import Create from '../Pay/Pay';
const Payment: FC<{}> = () => {
  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`Payment`}
      >
      </Header>
      <Content>
        <Grid container spacing = {3} >
          <Grid item xs = {12} sm = {12}  >
            <Typography variant="h3" gutterBottom align = "center" >
              แจ้งชำระเงิน
              <br/>
              </Typography>
              <Grid item xs = {12} sm = {12}  >
            <Create/>
            </Grid>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );
};

export default Payment;
