import React from 'react';
import './App.css';

import LeftDrawer from './Components/Drawer/LeftDrawer'

import Grid from '@material-ui/core/Grid';

import {MuiThemeProvider} from '@material-ui/core/styles';

import useMediaQuery from '@material-ui/core/useMediaQuery';
import { createMuiTheme, ThemeProvider } from '@material-ui/core/styles';

const theme = createMuiTheme({
  palette: {
    type: "dark",
    grey: {
      800: "#424242", // overrides failed
      900: "#121212" // overrides success
    },
    background: {
      paper: "#424242",
      default: "#303030"
    }
  }
});

function App() {


  return (
    <MuiThemeProvider theme={theme}>
      <Grid>
        <LeftDrawer></LeftDrawer>
      </Grid>
    </MuiThemeProvider>
  );
}

export default App;
