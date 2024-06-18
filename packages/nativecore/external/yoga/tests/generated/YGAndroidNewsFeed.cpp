/*
 * Copyright (c) Meta Platforms, Inc. and affiliates.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * clang-format off
 * @generated SignedSource<<8cfd20a845597bba085375ad4308cdc0>>
 * generated by gentest/gentest-driver.ts from gentest/fixtures/YGAndroidNewsFeed.html
 */

#include <gtest/gtest.h>
#include <yoga/Yoga.h>

TEST(YogaTest, android_news_feed) {
  const YGConfigRef config = YGConfigNew();

  const YGNodeRef root = YGNodeNewWithConfig(config);
  YGNodeStyleSetAlignContent(root, YGAlignStretch);
  YGNodeStyleSetPositionType(root, YGPositionTypeAbsolute);
  YGNodeStyleSetWidth(root, 1080);

  const YGNodeRef root_child0 = YGNodeNewWithConfig(config);
  YGNodeInsertChild(root, root_child0, 0);

  const YGNodeRef root_child0_child0 = YGNodeNewWithConfig(config);
  YGNodeStyleSetAlignContent(root_child0_child0, YGAlignStretch);
  YGNodeInsertChild(root_child0, root_child0_child0, 0);

  const YGNodeRef root_child0_child0_child0 = YGNodeNewWithConfig(config);
  YGNodeStyleSetAlignContent(root_child0_child0_child0, YGAlignStretch);
  YGNodeInsertChild(root_child0_child0, root_child0_child0_child0, 0);

  const YGNodeRef root_child0_child0_child0_child0 = YGNodeNewWithConfig(config);
  YGNodeStyleSetFlexDirection(root_child0_child0_child0_child0, YGFlexDirectionRow);
  YGNodeStyleSetAlignContent(root_child0_child0_child0_child0, YGAlignStretch);
  YGNodeStyleSetAlignItems(root_child0_child0_child0_child0, YGAlignFlexStart);
  YGNodeStyleSetMargin(root_child0_child0_child0_child0, YGEdgeStart, 36);
  YGNodeStyleSetMargin(root_child0_child0_child0_child0, YGEdgeTop, 24);
  YGNodeInsertChild(root_child0_child0_child0, root_child0_child0_child0_child0, 0);

  const YGNodeRef root_child0_child0_child0_child0_child0 = YGNodeNewWithConfig(config);
  YGNodeStyleSetFlexDirection(root_child0_child0_child0_child0_child0, YGFlexDirectionRow);
  YGNodeStyleSetAlignContent(root_child0_child0_child0_child0_child0, YGAlignStretch);
  YGNodeInsertChild(root_child0_child0_child0_child0, root_child0_child0_child0_child0_child0, 0);

  const YGNodeRef root_child0_child0_child0_child0_child0_child0 = YGNodeNewWithConfig(config);
  YGNodeStyleSetAlignContent(root_child0_child0_child0_child0_child0_child0, YGAlignStretch);
  YGNodeStyleSetWidth(root_child0_child0_child0_child0_child0_child0, 120);
  YGNodeStyleSetHeight(root_child0_child0_child0_child0_child0_child0, 120);
  YGNodeInsertChild(root_child0_child0_child0_child0_child0, root_child0_child0_child0_child0_child0_child0, 0);

  const YGNodeRef root_child0_child0_child0_child0_child1 = YGNodeNewWithConfig(config);
  YGNodeStyleSetAlignContent(root_child0_child0_child0_child0_child1, YGAlignStretch);
  YGNodeStyleSetFlexShrink(root_child0_child0_child0_child0_child1, 1);
  YGNodeStyleSetMargin(root_child0_child0_child0_child0_child1, YGEdgeRight, 36);
  YGNodeStyleSetPadding(root_child0_child0_child0_child0_child1, YGEdgeLeft, 36);
  YGNodeStyleSetPadding(root_child0_child0_child0_child0_child1, YGEdgeTop, 21);
  YGNodeStyleSetPadding(root_child0_child0_child0_child0_child1, YGEdgeRight, 36);
  YGNodeStyleSetPadding(root_child0_child0_child0_child0_child1, YGEdgeBottom, 18);
  YGNodeInsertChild(root_child0_child0_child0_child0, root_child0_child0_child0_child0_child1, 1);

  const YGNodeRef root_child0_child0_child0_child0_child1_child0 = YGNodeNewWithConfig(config);
  YGNodeStyleSetFlexDirection(root_child0_child0_child0_child0_child1_child0, YGFlexDirectionRow);
  YGNodeStyleSetAlignContent(root_child0_child0_child0_child0_child1_child0, YGAlignStretch);
  YGNodeStyleSetFlexShrink(root_child0_child0_child0_child0_child1_child0, 1);
  YGNodeInsertChild(root_child0_child0_child0_child0_child1, root_child0_child0_child0_child0_child1_child0, 0);

  const YGNodeRef root_child0_child0_child0_child0_child1_child1 = YGNodeNewWithConfig(config);
  YGNodeStyleSetAlignContent(root_child0_child0_child0_child0_child1_child1, YGAlignStretch);
  YGNodeStyleSetFlexShrink(root_child0_child0_child0_child0_child1_child1, 1);
  YGNodeInsertChild(root_child0_child0_child0_child0_child1, root_child0_child0_child0_child0_child1_child1, 1);

  const YGNodeRef root_child0_child0_child1 = YGNodeNewWithConfig(config);
  YGNodeStyleSetAlignContent(root_child0_child0_child1, YGAlignStretch);
  YGNodeInsertChild(root_child0_child0, root_child0_child0_child1, 1);

  const YGNodeRef root_child0_child0_child1_child0 = YGNodeNewWithConfig(config);
  YGNodeStyleSetFlexDirection(root_child0_child0_child1_child0, YGFlexDirectionRow);
  YGNodeStyleSetAlignContent(root_child0_child0_child1_child0, YGAlignStretch);
  YGNodeStyleSetAlignItems(root_child0_child0_child1_child0, YGAlignFlexStart);
  YGNodeStyleSetMargin(root_child0_child0_child1_child0, YGEdgeStart, 174);
  YGNodeStyleSetMargin(root_child0_child0_child1_child0, YGEdgeTop, 24);
  YGNodeInsertChild(root_child0_child0_child1, root_child0_child0_child1_child0, 0);

  const YGNodeRef root_child0_child0_child1_child0_child0 = YGNodeNewWithConfig(config);
  YGNodeStyleSetFlexDirection(root_child0_child0_child1_child0_child0, YGFlexDirectionRow);
  YGNodeStyleSetAlignContent(root_child0_child0_child1_child0_child0, YGAlignStretch);
  YGNodeInsertChild(root_child0_child0_child1_child0, root_child0_child0_child1_child0_child0, 0);

  const YGNodeRef root_child0_child0_child1_child0_child0_child0 = YGNodeNewWithConfig(config);
  YGNodeStyleSetAlignContent(root_child0_child0_child1_child0_child0_child0, YGAlignStretch);
  YGNodeStyleSetWidth(root_child0_child0_child1_child0_child0_child0, 72);
  YGNodeStyleSetHeight(root_child0_child0_child1_child0_child0_child0, 72);
  YGNodeInsertChild(root_child0_child0_child1_child0_child0, root_child0_child0_child1_child0_child0_child0, 0);

  const YGNodeRef root_child0_child0_child1_child0_child1 = YGNodeNewWithConfig(config);
  YGNodeStyleSetAlignContent(root_child0_child0_child1_child0_child1, YGAlignStretch);
  YGNodeStyleSetFlexShrink(root_child0_child0_child1_child0_child1, 1);
  YGNodeStyleSetMargin(root_child0_child0_child1_child0_child1, YGEdgeRight, 36);
  YGNodeStyleSetPadding(root_child0_child0_child1_child0_child1, YGEdgeLeft, 36);
  YGNodeStyleSetPadding(root_child0_child0_child1_child0_child1, YGEdgeTop, 21);
  YGNodeStyleSetPadding(root_child0_child0_child1_child0_child1, YGEdgeRight, 36);
  YGNodeStyleSetPadding(root_child0_child0_child1_child0_child1, YGEdgeBottom, 18);
  YGNodeInsertChild(root_child0_child0_child1_child0, root_child0_child0_child1_child0_child1, 1);

  const YGNodeRef root_child0_child0_child1_child0_child1_child0 = YGNodeNewWithConfig(config);
  YGNodeStyleSetFlexDirection(root_child0_child0_child1_child0_child1_child0, YGFlexDirectionRow);
  YGNodeStyleSetAlignContent(root_child0_child0_child1_child0_child1_child0, YGAlignStretch);
  YGNodeStyleSetFlexShrink(root_child0_child0_child1_child0_child1_child0, 1);
  YGNodeInsertChild(root_child0_child0_child1_child0_child1, root_child0_child0_child1_child0_child1_child0, 0);

  const YGNodeRef root_child0_child0_child1_child0_child1_child1 = YGNodeNewWithConfig(config);
  YGNodeStyleSetAlignContent(root_child0_child0_child1_child0_child1_child1, YGAlignStretch);
  YGNodeStyleSetFlexShrink(root_child0_child0_child1_child0_child1_child1, 1);
  YGNodeInsertChild(root_child0_child0_child1_child0_child1, root_child0_child0_child1_child0_child1_child1, 1);
  YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionLTR);

  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetLeft(root));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetTop(root));
  ASSERT_FLOAT_EQ(1080, YGNodeLayoutGetWidth(root));
  ASSERT_FLOAT_EQ(240, YGNodeLayoutGetHeight(root));

  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetLeft(root_child0));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetTop(root_child0));
  ASSERT_FLOAT_EQ(1080, YGNodeLayoutGetWidth(root_child0));
  ASSERT_FLOAT_EQ(240, YGNodeLayoutGetHeight(root_child0));

  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetLeft(root_child0_child0));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetTop(root_child0_child0));
  ASSERT_FLOAT_EQ(1080, YGNodeLayoutGetWidth(root_child0_child0));
  ASSERT_FLOAT_EQ(240, YGNodeLayoutGetHeight(root_child0_child0));

  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetLeft(root_child0_child0_child0));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetTop(root_child0_child0_child0));
  ASSERT_FLOAT_EQ(1080, YGNodeLayoutGetWidth(root_child0_child0_child0));
  ASSERT_FLOAT_EQ(144, YGNodeLayoutGetHeight(root_child0_child0_child0));

  ASSERT_FLOAT_EQ(36, YGNodeLayoutGetLeft(root_child0_child0_child0_child0));
  ASSERT_FLOAT_EQ(24, YGNodeLayoutGetTop(root_child0_child0_child0_child0));
  ASSERT_FLOAT_EQ(1044, YGNodeLayoutGetWidth(root_child0_child0_child0_child0));
  ASSERT_FLOAT_EQ(120, YGNodeLayoutGetHeight(root_child0_child0_child0_child0));

  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetLeft(root_child0_child0_child0_child0_child0));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetTop(root_child0_child0_child0_child0_child0));
  ASSERT_FLOAT_EQ(120, YGNodeLayoutGetWidth(root_child0_child0_child0_child0_child0));
  ASSERT_FLOAT_EQ(120, YGNodeLayoutGetHeight(root_child0_child0_child0_child0_child0));

  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetLeft(root_child0_child0_child0_child0_child0_child0));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetTop(root_child0_child0_child0_child0_child0_child0));
  ASSERT_FLOAT_EQ(120, YGNodeLayoutGetWidth(root_child0_child0_child0_child0_child0_child0));
  ASSERT_FLOAT_EQ(120, YGNodeLayoutGetHeight(root_child0_child0_child0_child0_child0_child0));

  ASSERT_FLOAT_EQ(120, YGNodeLayoutGetLeft(root_child0_child0_child0_child0_child1));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetTop(root_child0_child0_child0_child0_child1));
  ASSERT_FLOAT_EQ(72, YGNodeLayoutGetWidth(root_child0_child0_child0_child0_child1));
  ASSERT_FLOAT_EQ(39, YGNodeLayoutGetHeight(root_child0_child0_child0_child0_child1));

  ASSERT_FLOAT_EQ(36, YGNodeLayoutGetLeft(root_child0_child0_child0_child0_child1_child0));
  ASSERT_FLOAT_EQ(21, YGNodeLayoutGetTop(root_child0_child0_child0_child0_child1_child0));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetWidth(root_child0_child0_child0_child0_child1_child0));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetHeight(root_child0_child0_child0_child0_child1_child0));

  ASSERT_FLOAT_EQ(36, YGNodeLayoutGetLeft(root_child0_child0_child0_child0_child1_child1));
  ASSERT_FLOAT_EQ(21, YGNodeLayoutGetTop(root_child0_child0_child0_child0_child1_child1));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetWidth(root_child0_child0_child0_child0_child1_child1));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetHeight(root_child0_child0_child0_child0_child1_child1));

  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetLeft(root_child0_child0_child1));
  ASSERT_FLOAT_EQ(144, YGNodeLayoutGetTop(root_child0_child0_child1));
  ASSERT_FLOAT_EQ(1080, YGNodeLayoutGetWidth(root_child0_child0_child1));
  ASSERT_FLOAT_EQ(96, YGNodeLayoutGetHeight(root_child0_child0_child1));

  ASSERT_FLOAT_EQ(174, YGNodeLayoutGetLeft(root_child0_child0_child1_child0));
  ASSERT_FLOAT_EQ(24, YGNodeLayoutGetTop(root_child0_child0_child1_child0));
  ASSERT_FLOAT_EQ(906, YGNodeLayoutGetWidth(root_child0_child0_child1_child0));
  ASSERT_FLOAT_EQ(72, YGNodeLayoutGetHeight(root_child0_child0_child1_child0));

  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetLeft(root_child0_child0_child1_child0_child0));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetTop(root_child0_child0_child1_child0_child0));
  ASSERT_FLOAT_EQ(72, YGNodeLayoutGetWidth(root_child0_child0_child1_child0_child0));
  ASSERT_FLOAT_EQ(72, YGNodeLayoutGetHeight(root_child0_child0_child1_child0_child0));

  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetLeft(root_child0_child0_child1_child0_child0_child0));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetTop(root_child0_child0_child1_child0_child0_child0));
  ASSERT_FLOAT_EQ(72, YGNodeLayoutGetWidth(root_child0_child0_child1_child0_child0_child0));
  ASSERT_FLOAT_EQ(72, YGNodeLayoutGetHeight(root_child0_child0_child1_child0_child0_child0));

  ASSERT_FLOAT_EQ(72, YGNodeLayoutGetLeft(root_child0_child0_child1_child0_child1));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetTop(root_child0_child0_child1_child0_child1));
  ASSERT_FLOAT_EQ(72, YGNodeLayoutGetWidth(root_child0_child0_child1_child0_child1));
  ASSERT_FLOAT_EQ(39, YGNodeLayoutGetHeight(root_child0_child0_child1_child0_child1));

  ASSERT_FLOAT_EQ(36, YGNodeLayoutGetLeft(root_child0_child0_child1_child0_child1_child0));
  ASSERT_FLOAT_EQ(21, YGNodeLayoutGetTop(root_child0_child0_child1_child0_child1_child0));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetWidth(root_child0_child0_child1_child0_child1_child0));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetHeight(root_child0_child0_child1_child0_child1_child0));

  ASSERT_FLOAT_EQ(36, YGNodeLayoutGetLeft(root_child0_child0_child1_child0_child1_child1));
  ASSERT_FLOAT_EQ(21, YGNodeLayoutGetTop(root_child0_child0_child1_child0_child1_child1));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetWidth(root_child0_child0_child1_child0_child1_child1));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetHeight(root_child0_child0_child1_child0_child1_child1));

  YGNodeCalculateLayout(root, YGUndefined, YGUndefined, YGDirectionRTL);

  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetLeft(root));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetTop(root));
  ASSERT_FLOAT_EQ(1080, YGNodeLayoutGetWidth(root));
  ASSERT_FLOAT_EQ(240, YGNodeLayoutGetHeight(root));

  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetLeft(root_child0));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetTop(root_child0));
  ASSERT_FLOAT_EQ(1080, YGNodeLayoutGetWidth(root_child0));
  ASSERT_FLOAT_EQ(240, YGNodeLayoutGetHeight(root_child0));

  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetLeft(root_child0_child0));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetTop(root_child0_child0));
  ASSERT_FLOAT_EQ(1080, YGNodeLayoutGetWidth(root_child0_child0));
  ASSERT_FLOAT_EQ(240, YGNodeLayoutGetHeight(root_child0_child0));

  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetLeft(root_child0_child0_child0));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetTop(root_child0_child0_child0));
  ASSERT_FLOAT_EQ(1080, YGNodeLayoutGetWidth(root_child0_child0_child0));
  ASSERT_FLOAT_EQ(144, YGNodeLayoutGetHeight(root_child0_child0_child0));

  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetLeft(root_child0_child0_child0_child0));
  ASSERT_FLOAT_EQ(24, YGNodeLayoutGetTop(root_child0_child0_child0_child0));
  ASSERT_FLOAT_EQ(1044, YGNodeLayoutGetWidth(root_child0_child0_child0_child0));
  ASSERT_FLOAT_EQ(120, YGNodeLayoutGetHeight(root_child0_child0_child0_child0));

  ASSERT_FLOAT_EQ(924, YGNodeLayoutGetLeft(root_child0_child0_child0_child0_child0));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetTop(root_child0_child0_child0_child0_child0));
  ASSERT_FLOAT_EQ(120, YGNodeLayoutGetWidth(root_child0_child0_child0_child0_child0));
  ASSERT_FLOAT_EQ(120, YGNodeLayoutGetHeight(root_child0_child0_child0_child0_child0));

  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetLeft(root_child0_child0_child0_child0_child0_child0));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetTop(root_child0_child0_child0_child0_child0_child0));
  ASSERT_FLOAT_EQ(120, YGNodeLayoutGetWidth(root_child0_child0_child0_child0_child0_child0));
  ASSERT_FLOAT_EQ(120, YGNodeLayoutGetHeight(root_child0_child0_child0_child0_child0_child0));

  ASSERT_FLOAT_EQ(816, YGNodeLayoutGetLeft(root_child0_child0_child0_child0_child1));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetTop(root_child0_child0_child0_child0_child1));
  ASSERT_FLOAT_EQ(72, YGNodeLayoutGetWidth(root_child0_child0_child0_child0_child1));
  ASSERT_FLOAT_EQ(39, YGNodeLayoutGetHeight(root_child0_child0_child0_child0_child1));

  ASSERT_FLOAT_EQ(36, YGNodeLayoutGetLeft(root_child0_child0_child0_child0_child1_child0));
  ASSERT_FLOAT_EQ(21, YGNodeLayoutGetTop(root_child0_child0_child0_child0_child1_child0));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetWidth(root_child0_child0_child0_child0_child1_child0));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetHeight(root_child0_child0_child0_child0_child1_child0));

  ASSERT_FLOAT_EQ(36, YGNodeLayoutGetLeft(root_child0_child0_child0_child0_child1_child1));
  ASSERT_FLOAT_EQ(21, YGNodeLayoutGetTop(root_child0_child0_child0_child0_child1_child1));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetWidth(root_child0_child0_child0_child0_child1_child1));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetHeight(root_child0_child0_child0_child0_child1_child1));

  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetLeft(root_child0_child0_child1));
  ASSERT_FLOAT_EQ(144, YGNodeLayoutGetTop(root_child0_child0_child1));
  ASSERT_FLOAT_EQ(1080, YGNodeLayoutGetWidth(root_child0_child0_child1));
  ASSERT_FLOAT_EQ(96, YGNodeLayoutGetHeight(root_child0_child0_child1));

  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetLeft(root_child0_child0_child1_child0));
  ASSERT_FLOAT_EQ(24, YGNodeLayoutGetTop(root_child0_child0_child1_child0));
  ASSERT_FLOAT_EQ(906, YGNodeLayoutGetWidth(root_child0_child0_child1_child0));
  ASSERT_FLOAT_EQ(72, YGNodeLayoutGetHeight(root_child0_child0_child1_child0));

  ASSERT_FLOAT_EQ(834, YGNodeLayoutGetLeft(root_child0_child0_child1_child0_child0));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetTop(root_child0_child0_child1_child0_child0));
  ASSERT_FLOAT_EQ(72, YGNodeLayoutGetWidth(root_child0_child0_child1_child0_child0));
  ASSERT_FLOAT_EQ(72, YGNodeLayoutGetHeight(root_child0_child0_child1_child0_child0));

  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetLeft(root_child0_child0_child1_child0_child0_child0));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetTop(root_child0_child0_child1_child0_child0_child0));
  ASSERT_FLOAT_EQ(72, YGNodeLayoutGetWidth(root_child0_child0_child1_child0_child0_child0));
  ASSERT_FLOAT_EQ(72, YGNodeLayoutGetHeight(root_child0_child0_child1_child0_child0_child0));

  ASSERT_FLOAT_EQ(726, YGNodeLayoutGetLeft(root_child0_child0_child1_child0_child1));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetTop(root_child0_child0_child1_child0_child1));
  ASSERT_FLOAT_EQ(72, YGNodeLayoutGetWidth(root_child0_child0_child1_child0_child1));
  ASSERT_FLOAT_EQ(39, YGNodeLayoutGetHeight(root_child0_child0_child1_child0_child1));

  ASSERT_FLOAT_EQ(36, YGNodeLayoutGetLeft(root_child0_child0_child1_child0_child1_child0));
  ASSERT_FLOAT_EQ(21, YGNodeLayoutGetTop(root_child0_child0_child1_child0_child1_child0));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetWidth(root_child0_child0_child1_child0_child1_child0));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetHeight(root_child0_child0_child1_child0_child1_child0));

  ASSERT_FLOAT_EQ(36, YGNodeLayoutGetLeft(root_child0_child0_child1_child0_child1_child1));
  ASSERT_FLOAT_EQ(21, YGNodeLayoutGetTop(root_child0_child0_child1_child0_child1_child1));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetWidth(root_child0_child0_child1_child0_child1_child1));
  ASSERT_FLOAT_EQ(0, YGNodeLayoutGetHeight(root_child0_child0_child1_child0_child1_child1));

  YGNodeFreeRecursive(root);

  YGConfigFree(config);
}