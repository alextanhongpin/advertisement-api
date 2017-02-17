# Advertisement Microservice

__Work in Progress__

## Testing

To bootstrap testing:
```go
ginkgo bootstrap
```

To generate test specs:
```go
ginkgo generate users
```
# User Story

---

Part of the 'Campaign' Epic

@C01#Done
## Feature: Create Campaign

As a user,
I want to create a campaign,
To measure the success of my offerings.

@C02#Done
## Feature: View Campaign

As a user,
I want to view a campaign,
In order to get more info on the campaign

## Feature: Delete Campaign

As a user,
I want to delete a campaign,
In order to permanently remove it from my campaign list.

## Feature: Update Campaign

As a user,
I want to update my campaign,
So that others can view the campaign.

---

Part of the 'Advertisement' Epic.

@AD01
## Feature: View Ads

As a user,
I want to view ads,
In order to get informed on new deals.

@AD02
## Feature: Create Ads:

As a user,
I want to create ads,
In order to promote my offerings.

@AD03
## Feature: Rotate Ads

As a user,
I want to rotate my ads,
So that the user won't see similar results for the same deal.

@AD04
## Feature: Range Date Ads

As a user,
I want to show users specific ads at specific time,
To match their timezone.

@AD05
## Feature: Analytics

As a user,
I want to view analytics for the ads I created,
In order to measure the success metrics.

--- 
Part of the 'Miscellaneous' Epic

@MSC01#Done
## Feature: Serve static files

As a dev,
I want to serve static files,
So that I can use it on the client-side.


HTTP Verb   Path    Controller#Action   Used for
GET /photos photos#index    display a list of all photos
GET /photos/new photos#new  return an HTML form for creating a new photo
POST    /photos photos#create   create a new photo
GET /photos/:id photos#show display a specific photo
GET /photos/:id/edit    photos#edit return an HTML form for editing a photo
PATCH/PUT   /photos/:id photos#update   update a specific photo
DELETE  /photos/:id photos#destroy  delete a specific photo


All() Get the items for a resource
One() Get an item for a resource
Create() Create a new resource
Update() Update the resource
Delete() Delete a resource

