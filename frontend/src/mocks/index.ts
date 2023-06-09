const mockSite = {
  users: [
    'e02fc413-b7b8-4f92-858b-084e20b4e335',
    'b9f01df2-0ff9-4732-ada4-a49135fd0cc8',
  ],
  usersUiData: [
    {
      id: 'e02fc413-b7b8-4f92-858b-084e20b4e335',
      userId: 'e02fc413-b7b8-4f92-858b-084e20b4e335',
      sitename: 'Test',
      homepageImgUrl: '/heroBackground/default_1.jpeg',
      homepageText: 'Welcome!',
      homepageTextColor: '#ffffff',
    },
  ],
}
function getDefaultUiData(id: string, userId: string) {
  return { ...mockSite.usersUiData[0], id: id, userId: userId }
}
const mock = {}

export { mockSite, getDefaultUiData }
