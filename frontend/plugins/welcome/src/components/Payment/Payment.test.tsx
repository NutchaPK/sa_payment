import React from 'react';
import { render } from '@testing-library/react';
import Payment from './Payment';
import { ThemeProvider } from '@material-ui/core';
import { lightTheme } from '@backstage/theme';

describe('Payment', () => {
  it('should render', () => {
    const rendered = render(
      <ThemeProvider theme={lightTheme}>
        <Payment />
      </ThemeProvider>,
    );
    expect(rendered.baseElement).toBeInTheDocument();
  });
});
