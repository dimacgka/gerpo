package sql

//type StringJoinBuilder struct {
//	ctx context.Context
//	sql string
//}
//
//func (b *StringGroupBuilder) GroupBy(cols ...types.Column) {
//	for _, col := range cols {
//		if !col.IsAllowedAction(types.SQLActionGroup) {
//			continue
//			//TODO: log
//		}
//		b.sql += col.ToSQL(b.ctx)
//	}
//}