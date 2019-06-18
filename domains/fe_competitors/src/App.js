import React from 'react';
import clsx from 'clsx';

import { loadCSS } from 'fg-loadcss';
import AppBar from '@material-ui/core/AppBar';
import CssBaseline from '@material-ui/core/CssBaseline';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import Icon from '@material-ui/core/Icon';

import ResultContainer from './js/components/container/ResultContainer'

import './App.css';


const useStyles = makeStyles(theme => ({
  icon: {
    marginRight: theme.spacing(2),
  },
  heroContent: {
    backgroundColor: theme.palette.background.paper,
    padding: theme.spacing(8, 0, 6),
  },
  heroButtons: {
    marginTop: theme.spacing(4),
  },
  footer: {
    backgroundColor: theme.palette.background.paper,
    padding: theme.spacing(6),
  },
}));


function App() {
  const classes = useStyles();


    React.useEffect(() => {
        loadCSS(
            'https://use.fontawesome.com/releases/v5.1.0/css/all.css',
            document.querySelector('#font-awesome-css'),
        );
    }, []);

  return (
      <React.Fragment>
        <CssBaseline />
        <AppBar position="relative">
          <Toolbar>
            <Icon className={clsx(classes.icon, "fas fa-award")}></Icon>
            <Typography variant="h6" color="inherit" noWrap>
              Ribbonwall
            </Typography>
          </Toolbar>
        </AppBar>
        <main>
          {/* Hero unit */}
          <div className={classes.heroContent}>
            <Container maxWidth="sm">
              <Typography component="h1" variant="h2" align="center" color="textPrimary" gutterBottom>
                Ribbonwall
              </Typography>
              <Typography variant="h5" align="center" color="textSecondary" paragraph>
                Canadian Pony Club Results Tracker
              </Typography>
            </Container>
          </div>
          <ResultContainer></ResultContainer>
        </main>
        {/* Footer */}
        <footer className={classes.footer}>
          <Typography variant="h6" align="center" gutterBottom>
            Footer
          </Typography>
          <Typography variant="subtitle1" align="center" color="textSecondary" component="p">
            Put something here to give the footer a purpose!
          </Typography>
          <MadeWithLove />
        </footer>
        {/* End footer */}
      </React.Fragment>
  );
}

function MadeWithLove() {
  return (
      <Typography variant="body2" color="textSecondary" align="center">
        {'Built with love by John Finnson '}
      </Typography>
  );
}

export default App;
