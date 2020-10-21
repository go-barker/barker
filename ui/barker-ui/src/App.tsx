import React from 'react';
import { BotsListPage } from './bots/BotsListPage';
import { HashRouter, Switch, Route } from 'react-router-dom';
import BotPage from './bots/BotPage';
import { BotLoader, NewBotLoader } from './bots/BotLoader';
import { BotsListLoader } from './bots/BotsListLoader';

function App() {
    return (
        <HashRouter>
            <Switch>
                <Route exact path="/">
                    <BotsListLoader
                        render={(props) => <BotsListPage {...props} />}
                    />
                </Route>
                <Route exact path="/bots/new">
                    <NewBotLoader
                        render={(props) => (
                            <BotPage tab="edit" isNew {...props} />
                        )}
                    />
                </Route>
                <Route exact path="/bots/:botID">
                    <BotLoader
                        render={(props) => <BotPage tab="edit" {...props} />}
                    />
                </Route>
                <Route exact path="/bots/:botID/users">
                    <BotLoader
                        render={(props) => <BotPage tab="users" {...props} />}
                    />
                </Route>
                <Route exact path="/bots/:botID/campaigns">
                    <BotLoader
                        render={(props) => (
                            <BotPage tab="campaigns" {...props} />
                        )}
                    />
                </Route>
            </Switch>
        </HashRouter>
    );
}

export default App;
