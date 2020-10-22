import React, { FC } from 'react';
import { Typography,Grid, } from '@material-ui/core';
import { Content, Page, pageTheme, } from '@backstage/core';

const Home: FC<{}> = () => {
  return (
    <Page theme={pageTheme.home}>
      <Content>
        <Grid container justify="center">
          <Grid item xs = {6} sm = {3} justify="space-evenly" >
            <Typography variant="h6" gutterBottom >
              ระบบหลัก
              </Typography>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );
};
export default Home;
