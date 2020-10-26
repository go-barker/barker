import { Grid, makeStyles, Paper, TableCell } from '@material-ui/core';
import { Bot, PaginatorResponse } from 'barker-api';
import React, { FC } from 'react';
import { Link } from 'react-router-dom';
import { createListView } from '../createListView';
import BotsListAppBar from './BotsListAppBar';

const useStyles = makeStyles((theme) => ({
    paper: {
        width: '100%',
        padding: theme.spacing(2),
    },
}));

export interface BotsPageProps {
    items?: Bot[];
    error?: any;
    paging?: PaginatorResponse;
}

const BotsListView = createListView<Bot>({
    renderHeader: () => (
        <>
            <TableCell>ID</TableCell>
            <TableCell width="100%">Title</TableCell>
        </>
    ),
    renderRow: (item) => (
        <>
            <TableCell>{item.ID}</TableCell>
            <TableCell>
                <Link to={`/bots/${item.ID}`}>{item.Title}</Link>
            </TableCell>
        </>
    ),
});

export const BotsListPage: FC<BotsPageProps> = ({ items, error, paging }) => {
    const classes = useStyles();
    return (
        <Grid container>
            <BotsListAppBar />
            <Paper className={classes.paper}>
                <BotsListView items={items} error={error} paging={paging} />
            </Paper>
        </Grid>
    );
};
