import React, { Component } from "react";
import Container from "@material-ui/core/Container";
import Grid from "@material-ui/core/Grid";
import Card from "@material-ui/core/Card";
import CardMedia from "@material-ui/core/CardMedia";
import CardContent from "@material-ui/core/CardContent";
import Typography from "@material-ui/core/Typography";
import CardActions from "@material-ui/core/CardActions";
import Button from "@material-ui/core/Button";
import { withStyles } from "@material-ui/core";

const styles = {
    cardGrid: {
        // paddingTop: theme.spacing(8),
        // paddingBottom: theme.spacing(8),
    },
    card: {
        height: '100%',
        display: 'flex',
    },
    cardMedia: {
        height: 140,
        width: 140
    },
    cardContent: {
        display: 'flex',
        flexDirection: 'column',
    },
};


class ResultContainer extends Component {

    constructor(props) {
        super(props);
        this.state = {
            error: null,
            isLoaded: false,
            items: []
        };
    }

    componentDidMount() {
        fetch("http://localhost:8080/api/competition_results")
            .then(res => res.json())
            .then(
                (result) => {
                    this.setState({
                        isLoaded: true,
                        items: result
                    });
                },
                // Note: it's important to handle errors here
                // instead of a catch() block so that we don't swallow
                // exceptions from actual bugs in components.
                (error) => {
                    this.setState({
                        isLoaded: true,
                        error
                    });
                }
            )
    }

    render () {
        {
            const { classes } = this.props;
            const { error, isLoaded, items } = this.state;
            if (error) {
                return <div>Error: {error.message}</div>;
            } else if (!isLoaded) {
                return <div>Loading...</div>;
            } else {
                return (
                    <Container className={classes.cardGrid} maxWidth="md">
                    {/* End hero unit */}
                    <Grid container spacing={4}>
                        {items.map(card => (
                            <Grid item key={card.uuid} xs={12} sm={6} md={4}>
                                <Card className={classes.card}>
                                    <CardMedia
                                        className={classes.cardMedia}
                                        image="https://banner2.kisspng.com/20171127/ca8/red-rosette-ribbon-png-clipar-image-5a1be8acf191a4.0247494115117784769895.jpg"
                                        title="Image title"
                                    />
                                    <CardContent className={classes.cardContent}>
                                        <Typography gutterBottom variant="h5" component="h2">
                                            {card.competitor.firstName}
                                        </Typography>
                                        <Typography>
                                            Placing: {card.placing}
                                        </Typography>
                                    </CardContent>
                                    {/*<CardActions>*/}
                                    {/*    <Button size="small" color="primary">*/}
                                    {/*        View*/}
                                    {/*    </Button>*/}
                                    {/*    <Button size="small" color="primary">*/}
                                    {/*        Edit*/}
                                    {/*    </Button>*/}
                                    {/*</CardActions>*/}
                                </Card>
                            </Grid>
                        ))}
                    </Grid>
                </Container>)
            }
        }
    }
}

export default withStyles(styles)(ResultContainer);