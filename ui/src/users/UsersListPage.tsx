import { Grid, TableCell } from '@material-ui/core';
import { PaginatorResponse, User } from 'barker-api';
import React, { FC } from 'react';
import { createListView } from '../createListView';

export interface UsersListPageProps {
    items?: User[];
    error?: any;
    paging?: PaginatorResponse;
}

const UsersListView = createListView<User>({
    renderHeader: () => (
        <>
            <TableCell>TelegramID</TableCell>
            <TableCell width="100%">DisplayName</TableCell>
            <TableCell width="100%">UserName</TableCell>
            <TableCell width="100%">FirstName</TableCell>
            <TableCell width="100%">LastName</TableCell>
        </>
    ),
    renderRow: (item) => (
        <>
            <TableCell>{item.TelegramID}</TableCell>
            <TableCell>
                {/* <Link to={`/bots/${item.BotID}/users/${item.TelegramID}`}> */}
                {item.DisplayName || '<no>'}
                {/* </Link> */}
            </TableCell>
            <TableCell>
                {item.UserName ? (
                    <a
                        target="_blank"
                        rel="noopener noreferrer"
                        href={`https://t.me/${item.UserName}`}
                    >
                        @{item.UserName}
                    </a>
                ) : (
                    '<no>'
                )}
            </TableCell>
            <TableCell>{item.FirstName || '<no>'}</TableCell>
            <TableCell>{item.LastName || '<no>'}</TableCell>
        </>
    ),
});

export const UsersListPage: FC<UsersListPageProps> = ({
    items,
    error,
    paging,
}) => {
    return (
        <Grid container>
            <UsersListView items={items} error={error} paging={paging} />
        </Grid>
    );
};
