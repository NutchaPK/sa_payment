import React from 'react';
import PaymentIcon from '@material-ui/icons/Payment';
import HomeIcon from '@material-ui/icons/Home';
import SignOut from '@material-ui/icons/Settings';
import HistoryIcon from '@material-ui/icons/History';
import {
  Sidebar,
  SidebarItem,
  SidebarDivider,
  SidebarSpace,
  SidebarPinButton,
} from '@backstage/core';

export const AppSidebar = () => (
  <Sidebar>
    <SidebarDivider />
    {/* Global nav, not org-specific */}
    <SidebarItem icon= {HomeIcon} to="Home" text="Home" />
    <SidebarItem icon= {PaymentIcon} to="Payment" text="Payment" />
    <SidebarItem icon= {HistoryIcon} to="Table" text="History Payment" />
    {/* End global nav */}
    <SidebarDivider />
    <SidebarSpace />
    <SidebarDivider />
    <SidebarItem
      icon={SignOut}
      to="./"
      text="Sign Out"
    />
    {/* <SidebarUserSettings  /> */}
    <SidebarPinButton />
  </Sidebar>
);