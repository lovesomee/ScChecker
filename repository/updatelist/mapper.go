package updatelist

import (
	"sc-profile/models"
)

func UpdateListToDb(updateList models.UpdateList) DbUpdateList {
	return DbUpdateList{
		ItemId: updateList.ItemId,
	}
}

func DbUpdateListToUpdateList(dbUpdateList DbUpdateList) models.UpdateList {
	return models.UpdateList{
		ItemId: dbUpdateList.ItemId,
	}
}

func DbUpdateListsToUpdateLists(dbUpdateLists []DbUpdateList) []models.UpdateList {
	updateLists := make([]models.UpdateList, 0, len(dbUpdateLists))

	for _, updateList := range dbUpdateLists {
		updateLists = append(updateLists, DbUpdateListToUpdateList(updateList))
	}

	return updateLists
}
