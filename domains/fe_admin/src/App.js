// in src/App.js
import React, { Component } from 'react';
import {makeStyles, withStyles} from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import IconButton from '@material-ui/core/IconButton';
import MenuIcon from '@material-ui/icons/Menu';
import Container from "@material-ui/core/Container";
import Grid from "@material-ui/core/Grid";
import {Admin} from "react-admin";
import CompetitionResults from "./js/components/container";
import AdminBrowser from "./js/adminBrowser";

const styles = {
    root: {
        flexGrow: 1,
    },
    menuButton: {
        marginRight: 2,
    },
    title: {
        flexGrow: 1,
    },

    icon: {
        marginRight: 2,
    },
    heroContent: {
        padding: "64px 0px 48px"
    },
};

class App extends Component {
    login() {
        this.props.auth.login();
    }

    logout() {
        this.props.auth.logout();
    }

    componentDidMount() {
        const { renewSession } = this.props.auth;

        if (localStorage.getItem('isLoggedIn') === 'true') {
            renewSession();
        }
    }

    render() {
        const { push, classes, ...props } = this.props;
        const { isAuthenticated } = this.props.auth;
        return (
            <React.Fragment>
                {
                    isAuthenticated() && (
                        <AdminBrowser {...props}  > </AdminBrowser>
                    )
                }
                {
                    !isAuthenticated() && (
                        <React.Fragment>
                            <AppBar position="static">
                                <Toolbar>
                                    <IconButton edge="start" className={classes.menuButton} color="inherit" aria-label="Menu">
                                        <MenuIcon />
                                    </IconButton>
                                    <Typography variant="h6" className={classes.title}>
                                        Ribbonwall Admin
                                    </Typography>
                                    <Button color="inherit" onClick={this.login.bind(this)}>Login</Button>
                                </Toolbar>
                            </AppBar>
                            <main>
                                    {/* Hero unit */}
                                <div className={classes.heroContent}>
                                    <Container maxWidth="sm">
                                        <Typography component="h1" variant="h2" align="center" color="textPrimary" gutterBottom>
                                            Ribbonwall Admin
                                        </Typography>
                                        <Typography variant="h5" align="center" color="textSecondary" paragraph>
                                            This is a discription for the Ribbonwall admin. You are seeing this because you have
                                            to log in.
                                        </Typography>
                                        <div className={classes.heroButtons}>
                                            <Grid container spacing={2} justify="center">
                                                <Button variant="contained" color="primary" onClick={this.login.bind(this)}>Login</Button>
                                            </Grid>
                                        </div>
                                    </Container>
                                </div>
                            </main>
                        </React.Fragment>
                    )
                }
         </React.Fragment>
        );
    }
}

export default withStyles(styles)(App);