import { Grid, TableCell, Button } from '@material-ui/core';
import { Campaign, PaginatorResponse } from 'barker-api';
import React, { FC } from 'react';
import { Link, useParams } from 'react-router-dom';
import { createListView } from '../createListView';

export interface CampaignsListPageProps {
    items?: Campaign[];
    error?: any;
    paging?: PaginatorResponse;
}

const CampaignsListView = createListView<Campaign>({
    renderHeader: () => (
        <>
            <TableCell>ID</TableCell>
            <TableCell width="100%">Title</TableCell>
            <TableCell width="100%">Message</TableCell>
            <TableCell width="100%">Active</TableCell>
        </>
    ),
    renderRow: (item) => (
        <>
            <TableCell>{item.ID}</TableCell>
            <TableCell>
                <Link to={`/bots/${item.BotID}/campaigns/${item.ID}`}>
                    {item.Title || '<no>'}
                </Link>
            </TableCell>
            <TableCell>{item.Message || '<no>'}</TableCell>
            <TableCell>{item.Active ? 'Active' : 'Inactive'}</TableCell>
        </>
    ),
});

export const CampaignsListPage: FC<CampaignsListPageProps> = ({
    items,
    error,
    paging,
}) => {
    const { botID } = useParams<{ botID?: string }>();
    return (
        <Grid container>
            <Button
                color="secondary"
                variant="contained"
                component={Link}
                to={`/bots/${botID}/campaigns/new`}
            >
                Create campaign
            </Button>
            <CampaignsListView items={items} error={error} paging={paging} />
        </Grid>
    );
};
