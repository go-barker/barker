import React from 'react';
import BotsPage from './bots/BotsPage';
import { HashRouter, Switch, Route } from 'react-router-dom';
import BotPage from './bots/BotPage';
import { BotLoader, NewBotLoader } from './bots/BotLoader';

function App() {
    return (
        <HashRouter>
            <Switch>
                <Route exact path="/">
                    <BotsPage />
                </Route>
                <Route exact path="/bots/new">
                    <NewBotLoader
                        render={(props) => (
                            <BotPage tab="edit" isNew {...props} />
                        )}
                    />
                </Route>
                <Route exact path="/bots/:id">
                    <BotLoader
                        render={(props) => <BotPage tab="edit" {...props} />}
                    />
                </Route>
                <Route exact path="/bots/:id/users">
                    <BotLoader
                        render={(props) => <BotPage tab="users" {...props} />}
                    />
                </Route>
                <Route exact path="/bots/:id/campaigns">
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
