// in src/users.js
import React, {Fragment} from 'react';
import { List, Datagrid, TextField, EmailField } from 'react-admin';
import { Route, Switch } from 'react-router';
import { Link } from "react-router-dom";
import { Drawer, withStyles } from '@material-ui/core';
import { CardActions, CreateButton, ExportButton, RefreshButton, Button } from 'react-admin';
import UploadCompetitionResults from "./upload_competition_results";
import CloudUploadIcon from '@material-ui/icons/CloudUpload';

const styles = {
    drawerContent: {
        width: 300
    }
};

const CompetitionResultsListActions = ({
   bulkActions,
   basePath,
   currentSort,
   displayedFilters,
   exporter,
   filters,
   filterValues,
   onUnselectItems,
   resource,
   selectedIds,
   showFilter,
   total,
   routeChange
}) => (
    <CardActions>
        <Button color="primary" component={Link} to="/competition_results/upload" label={"Upload"}>
            <CloudUploadIcon/>
        </Button>
        <ExportButton
            disabled={total === 0}
            resource={resource}
            sort={currentSort}
            filter={filterValues}
            exporter={exporter}
        />
        <RefreshButton />
        {/*<CreateButton basePath={basePath} />*/}
    </CardActions>
);

class CompetitionResultsList extends React.Component {
    routeChange() {
        let path = `/competition_results/upload`;
        this.props.history.push(path);
    }
    render() {
        const { push, classes, ...props } = this.props;
        return (
            <Fragment>
                <List {...props} title="Competition Results"
                      actions={<CompetitionResultsListActions routeChange={this.routeChange} />}
                >
                    <Datagrid rowClick="edit">
                        <TextField source="uuid" />
                        <TextField source="placing" />
                    </Datagrid>
                </List>
                <Route path="/competition_results/upload">
                    {({ match }) => (
                        <Drawer
                            open={!!match}
                            anchor="right"
                            onClose={this.handleClose}
                        >
                            <UploadCompetitionResults
                                className={classes.drawerContent}
                                // onCancel={this.handleClose}
                                {...props}
                            />
                        </Drawer>
                    )}
                </Route>
            </Fragment>
        );
    }

    // handleClose = () => {
    //     this.props.push('/competition_results');
    // };
}

export default withStyles(styles)(CompetitionResultsList);
